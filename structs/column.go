package structs

import (
	"strings"
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

var dbTypeToGoTypes = map[string]string{
	"tinyint":            "int8",
	"smallint":           "int16",
	"mediumint":          "int",
	"int":                "int",
	"bigint":             "int64",
	"tinyint unsigned":   "uint8",
	"smallint unsigned":  "uint16",
	"mediumint unsigned": "uint",
	"int unsigned":       "uint",
	"bigint unsigned":    "uint64",
	"float":              "float32",
	"double":             "float64",
}

// GoType returns type in go of the column
func (c Column) GoType() string {
	if c.FullDBType == "tinyint(1)" {
		return "bool"
	}

	for dbType, goType := range dbTypeToGoTypes {
		if c.DBType == dbType || strings.HasPrefix(c.DBType, dbType+"(") {
			return goType
		}
	}

	if strings.HasPrefix(c.DBType, "varchar") || strings.HasPrefix(c.DBType, "char") {
		if c.Nullable {
			return "null.String"
		}
		return "string"
	}

	if strings.HasPrefix(c.DBType, "varbinary") || strings.HasPrefix(c.DBType, "binary") {
		return "[]byte"
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

	panic("Unsupported db type: " + c.DBType)
}

// GoTypeNotNull returns type in go of the column as it's not nullable
func (c Column) GoTypeNotNull() string {
	t := c.GoType()
	if t == "null.String" {
		return "string"
	}

	return t
}

// GoName returns the variable name for go of the column
func (c Column) GoName() string {
	return getGoName(c.Name)
}

// IsNullable returns if the column is nullable as string
func (c Column) IsNullable() string {
	if c.Nullable {
		return "true"
	}

	return "false"
}

// HasDefault returns if the column has default value
func (c Column) HasDefault() bool {
	if c.Default == "auto_increment" {
		return false
	}

	return len(c.Default) > 0
}

func getQuotedStr(str string) string {
	if strings.HasPrefix(str, "\"") && strings.HasSuffix(str, "\"") {
		return str
	}
	return "\"" + str + "\""
}

// GoDefaultValue returns the go value of column's default value
func (c Column) GoDefaultValue() string {
	goType := c.GoType()
	lowerCaseDefault := strings.ToLower(c.Default)
	if goType == "string" || c.IsEnum() {
		return getQuotedStr(lowerCaseDefault)
	}
	if goType == "string" || c.IsSet() {
		lowerCaseNoSpaceDefault := strings.ReplaceAll(lowerCaseDefault, " ", "")
		if strings.HasPrefix(lowerCaseNoSpaceDefault, "(") && strings.HasSuffix(lowerCaseNoSpaceDefault, ")") {
			return lowerCaseNoSpaceDefault[1 : len(lowerCaseNoSpaceDefault)-1]
		}
		return getQuotedStr(lowerCaseDefault)
	}

	if goType == "time.Time" {
		if strings.Contains(c.Default, "CURRENT_TIMESTAMP") {
			return "time.Now()"
		}
		return c.Default
	}

	if goType == "bool" {
		if c.Default == "0" {
			return "false"
		}
		return "true"
	}

	if strings.Contains(goType, "int") || strings.Contains(goType, "float") {
		return goType + "(" + strings.ReplaceAll(c.Default, `"`, "") + ")"
	}

	return goType + "(" + c.Default + ")"
}

// IsEnum returns if column type is enum
func (c Column) IsEnum() bool {
	return strings.HasPrefix(c.DBType, "enum")
}

// IsSet returns if column type is set
func (c Column) IsSet() bool {
	return strings.HasPrefix(c.DBType, "set")
}

// IsBool returns if column type is boolean as tinyint(1)
func (c Column) IsBool() bool {
	return c.FullDBType == "tinyint(1)"
}

// GetEnumConst returns enum const definitions in go
func (c Column) GetEnumConst() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = c.GoType() + getGoName(enum) + " " + c.GoType() + " = " + `"` + enum + `"`
	}

	return strings.Join(elements, "\n")
}

// GetSetConst returns set const definitions in go
func (c Column) GetSetConst() string {
	enums := strings.Split(strings.ReplaceAll(getValue(c.FullDBType), "'", ""), ",")
	elements := make([]string, len(enums))
	for i, enum := range enums {
		elements[i] = c.GoType() + getGoName(enum) + " " + c.GoType() + " = " + `"` + enum + `"`
	}

	return strings.Join(elements, "\n")
}

// IsPrimaryKey returns if column is primary key
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

// IsAutoIncrement returns if column value is auto incremented
func (c Column) IsAutoIncrement() bool {
	return c.Default == "auto_increment"
}
