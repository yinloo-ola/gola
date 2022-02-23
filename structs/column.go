package structs

import (
	"fmt"
	"strings"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/vitess/go/vt/sqlparser"
)

// modified from: /home/wuvist/go/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/drivers/column.go

// Column holds information about a database column.
// Types are Go types, converted by TranslateColumnType.
type Column struct {
	Name      string `json:"name" toml:"name"`
	Type      string `json:"type" toml:"type"`
	DBType    string `json:"db_type" toml:"db_type"`
	Default   string `json:"default" toml:"default"`
	Comment   string `json:"comment" toml:"comment"`
	Nullable  bool   `json:"nullable" toml:"nullable"`
	Unique    bool   `json:"unique" toml:"unique"`
	Validated bool   `json:"validated" toml:"validated"`

	// Postgres only extension bits
	// ArrType is the underlying data type of the Postgres
	// ARRAY type. See here:
	// https://www.postgresql.org/docs/9.1/static/infoschema-element-types.html
	ArrType *string `json:"arr_type" toml:"arr_type"`
	UDTName string  `json:"udt_name" toml:"udt_name"`
	// DomainName is the domain type name associated to the column. See here:
	// https://www.postgresql.org/docs/10/extend-type-system.html#EXTEND-TYPE-SYSTEM-DOMAINS
	DomainName *string `json:"domain_name" toml:"domain_name"`

	// MySQL only bits
	// Used to get full type, ex:
	// tinyint(1) instead of tinyint
	// Used for "tinyint-as-bool" flag
	FullDBType string `json:"full_db_type" toml:"full_db_type"`

	// MS SQL only bits
	// Used to indicate that the value
	// for this column is auto generated by database on insert (i.e. - timestamp (old) or rowversion (new))
	AutoGenerated bool `json:"auto_generated" toml:"auto_generated"`

	Table *Table
}

var dbTypeToSQLTypes = map[string]string{
	"tinyint":           "sql.Int8",
	"smallint":          "sql.Int16",
	"int":               "sql.Int32",
	"bigint":            "sql.Int64",
	"tinyint unsigned":  "sql.Uint8",
	"smallint unsigned": "sql.Uint16",
	"int unsigned":      "sql.Uint32",
	"bigint unsigned":   "sql.Uint64",
}

var dbTypeToGoTypes = map[string]string{
	"tinyint":           "int8",
	"smallint":          "int16",
	"int":               "int",
	"bigint":            "int64",
	"tinyint unsigned":  "uint8",
	"smallint unsigned": "uint16",
	"int unsigned":      "uint",
	"bigint unsigned":   "uint64",
}

var dbTypeToPHPTypes = map[string]string{
	"tinyint":           "int",
	"smallint":          "int",
	"int":               "int",
	"bigint":            "int",
	"tinyint unsigned":  "int",
	"smallint unsigned": "int",
	"int unsigned":      "int",
	"bigint unsigned":   "int",
	"double":            "float",
}

func (c Column) SQLType() string {
	if sqlType, ok := dbTypeToSQLTypes[c.FullDBType]; ok {
		return sqlType
	}
	unsignedString := ""
	if strings.HasSuffix(c.FullDBType, "unsigned") {
		unsignedString = " unsigned"
	}
	if strings.HasPrefix(c.DBType, "tinyint") {
		if sqlType, ok := dbTypeToSQLTypes["tinyint"+unsignedString]; ok {
			return sqlType
		}
	} else if strings.HasPrefix(c.DBType, "smallint") {
		if sqlType, ok := dbTypeToSQLTypes["smallint"+unsignedString]; ok {
			return sqlType
		}
	} else if strings.HasPrefix(c.DBType, "bigint") {
		if sqlType, ok := dbTypeToSQLTypes["bigint"+unsignedString]; ok {
			return sqlType
		}
	} else if strings.HasPrefix(c.DBType, "int") {
		if sqlType, ok := dbTypeToSQLTypes["int"+unsignedString]; ok {
			return sqlType
		}
	}

	if strings.HasPrefix(c.DBType, "varchar") {
		size := getValue(c.FullDBType)

		return fmt.Sprintf("sql.MustCreateStringWithDefaults(sqltypes.VarChar, %s)", size)
	}

	if strings.HasPrefix(c.DBType, "decimal") || strings.HasPrefix(c.DBType, "float") {
		return "sql.Float32"
	}

	if strings.HasPrefix(c.DBType, "double") {
		return "sql.Float64"
	}

	if strings.HasPrefix(c.DBType, "enum") {
		enums := strings.ReplaceAll(getValue(c.FullDBType), "'", "\"")

		return fmt.Sprintf("sql.MustCreateEnumType([]string{%s}, sql.Collation_Default)", enums)
	}

	if strings.HasPrefix(c.DBType, "set") {
		vals := strings.ReplaceAll(getValue(c.FullDBType), "'", "\"")

		return fmt.Sprintf("sql.MustCreateSetType([]string{%s}, sql.Collation_Default)", vals)
	}

	if strings.HasPrefix(c.DBType, "char") {
		size := getValue(c.FullDBType)

		return fmt.Sprintf("sql.MustCreateStringWithDefaults(sqltypes.Char, %s)", size)
	}

	if strings.HasPrefix(c.DBType, "timestamp") {
		return "sql.Timestamp"
	}

	if strings.Contains(c.DBType, "text") {
		return "sql.Text"
	}

	if strings.Contains(c.DBType, "blob") {
		return "sql.Blob"
	}

	ct := &sqlparser.ColumnType{
		Type: c.DBType,
	}
	res, err := sql.ColumnTypeToType(ct)
	if err != nil {
		panic(err)
	}

	baseType := strings.ToLower(res.Type().String())
	return baseType
}

func (c Column) GoType() string {
	if goType, ok := dbTypeToGoTypes[c.DBType]; ok {
		return goType
	}

	if strings.HasPrefix(c.DBType, "varchar") || strings.HasPrefix(c.DBType, "char") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "decimal") {
		return "float32"
	}

	if c.IsEnum() {
		return c.Table.ClassName() + c.GoName()
	}

	if c.IsSet() {
		return c.Table.ClassName() + c.GoName()
	}

	if strings.HasPrefix(c.DBType, "set") {
		return "string"
	}

	if strings.Contains(c.DBType, "text") || strings.HasPrefix(c.DBType, "blob") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "timestamp") {
		return "time.Time"
	}

	ct := &sqlparser.ColumnType{
		Type: c.DBType,
	}
	res, err := sql.ColumnTypeToType(ct)
	if err != nil {
		panic(err)
	}

	baseType := strings.ToLower(res.Type().String())
	return baseType
}

func (c Column) PHPType() string {
	if goType, ok := dbTypeToPHPTypes[c.DBType]; ok {
		return goType
	}

	if strings.HasPrefix(c.DBType, "varchar") || strings.HasPrefix(c.DBType, "char") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "decimal") {
		return "float"
	}

	if c.IsEnum() {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "set") {
		return "string"
	}

	if strings.Contains(c.DBType, "text") || strings.HasPrefix(c.DBType, "blob") {
		return "string"
	}

	if strings.HasPrefix(c.DBType, "timestamp") {
		return "time.Time"
	}

	ct := &sqlparser.ColumnType{
		Type: c.DBType,
	}
	res, err := sql.ColumnTypeToType(ct)
	if err != nil {
		panic(err)
	}

	baseType := strings.ToLower(res.Type().String())
	return baseType
}

func (c Column) GoName() string {
	return getGoName(c.Name)
}

func (c Column) IsNullable() string {
	if c.Nullable {
		return "true"
	}

	return "false"
}

func (c Column) HasDefault() bool {
	if c.Default == "auto_increment" {
		return false
	}

	return len(c.Default) > 0
}

func (c Column) GoDefaultValue() string {
	goType := c.GoType()
	lowerCaseDefault := strings.ToLower(c.Default)
	if goType == "string" || c.IsEnum() {
		if strings.HasPrefix(lowerCaseDefault, "\"") && strings.HasSuffix(lowerCaseDefault, "\"") {
			return lowerCaseDefault
		}
		return "\"" + lowerCaseDefault + "\""
	}
	if goType == "string" || c.IsSet() {
		lowerCaseNoSpaceDefault := strings.ReplaceAll(lowerCaseDefault, " ", "")
		if strings.HasPrefix(lowerCaseNoSpaceDefault, "(") && strings.HasSuffix(lowerCaseNoSpaceDefault, ")") {
			return lowerCaseNoSpaceDefault[1 : len(lowerCaseNoSpaceDefault)-1]
		}
		return lowerCaseNoSpaceDefault
	}

	if goType == "time.Time" {
		if strings.Contains(c.Default, "CURRENT_TIMESTAMP") {
			return "time.Now()"
		}
		return c.Default
	}

	if strings.Contains(goType, "int") || strings.Contains(goType, "float") {
		return goType + "(" + strings.ReplaceAll(c.Default, `"`, "") + ")"
	}

	return goType + "(" + c.Default + ")"
}

func (c Column) GetColumnDefault() string {
	if !c.HasDefault() {
		return ""
	}

	return ", Default: default" + c.GoName()
}

func (c Column) IsEnum() bool {
	return strings.HasPrefix(c.DBType, "enum")
}
func (c Column) IsSet() bool {
	return strings.HasPrefix(c.DBType, "set")
}

func (c Column) GetEnumConst() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = c.GoType() + getGoName(enum) + " " + c.GoType() + " = " + `"` + enum + `"`
	}

	return strings.Join(elements, "\n")
}

func (c Column) GetSetConst() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = c.GoType() + getGoName(enum) + " " + c.GoType() + " = " + `"` + enum + `"`
	}

	return strings.Join(elements, "\n")
}

func (c Column) IsPrimaryKey() bool {
	if c.Table.PKey == nil {
		return false
	}

	for _, pc := range c.Table.PKey.Columns {
		if pc == c.Name {
			return true
		}
	}

	return false
}

func (c Column) IsAutoIncrement() bool {
	if c.Default == "auto_increment" {
		return true
	}

	return false
}
