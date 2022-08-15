// Code generated by gola 0.0.3; DO NOT EDIT.

package account

import (
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"
)

var DBName string = "testdata"

func Setup(dbname string) {
	DBName = dbname
}

// Account represents account table
type Account struct {
	//  int
	UserId
	// user account type enum('free','vip')
	Type
	// user country code mediumint unsigned
	CountryCode
	// Account money int
	Money
}

type AccountTable struct{}

func (*AccountTable) GetTableName() string {
	return "account"
}

var table *AccountTable

type PK struct {
	UserId      int
	CountryCode uint
}

type withPK interface {
	GetUserId() int
	GetCountryCode() uint
}

// FetchAccountByPKs returns a row from account table with given primary key value
func FetchAccountByPK(val PK) *Account {
	return coredb.FetchByPK[Account](DBName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FetchByPKs returns a row with selected fields from account table with given primary key value
func FetchByPK[T any](val PK) *T {
	return coredb.FetchByPK[T](DBName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FindOneAccount returns a row from account table with arbitary where query
// whereSQL must start with "where ..."
func FindOneAccount(whereSQL string, params ...any) *Account {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[Account](w, DBName)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt("SELECT COUNT(*) FROM `account` "+whereSQL, DBName, params...)
}

// FindOne returns a row with selected fields from account table with arbitary where query
// whereSQL must start with "where ..."
func FindOne[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](w, DBName)
}

// FindAccount returns rows from account table with arbitary where query
// whereSQL must start with "where ..."
func FindAccount(whereSQL string, params ...any) ([]*Account, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[Account](w, DBName)
}

// Find returns rows with selected fields from account table with arbitary where query
// whereSQL must start with "where ..."
func Find[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](w, DBName)
}

// Column types
type AccountType string

const (
	AccountTypeFree AccountType = "free"
	AccountTypeVip  AccountType = "vip"
)

// UserId field
//
type UserId struct {
	val int
}

func (c *UserId) GetUserId() int {
	return c.val
}

func (c *UserId) GetColumnName() string {
	return "user_id"
}

func (c *UserId) IsPrimaryKey() bool {
	return true
}

func (c *UserId) GetValPointer() any {
	return &c.val
}

func (c *UserId) GetTableType() coredb.TableType {
	return table
}

// Type field
// user account type
type Type struct {
	_updated bool
	val      AccountType
}

func (c *Type) GetType() AccountType {
	return c.val
}

func (c *Type) SetType(val AccountType) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Type) IsUpdated() bool {
	return c._updated
}

func (c *Type) resetUpdated() {
	c._updated = false
}

func (c *Type) GetColumnName() string {
	return "type"
}

func (c *Type) IsPrimaryKey() bool {
	return false
}

func (c *Type) GetValPointer() any {
	return &c.val
}

func (c *Type) GetTableType() coredb.TableType {
	return table
}

// CountryCode field
// user country code
type CountryCode struct {
	val uint
}

func (c *CountryCode) GetCountryCode() uint {
	return c.val
}

func (c *CountryCode) GetColumnName() string {
	return "country_code"
}

func (c *CountryCode) IsPrimaryKey() bool {
	return true
}

func (c *CountryCode) GetValPointer() any {
	return &c.val
}

func (c *CountryCode) GetTableType() coredb.TableType {
	return table
}

// Money field
// Account money
type Money struct {
	_updated bool
	val      int
}

func (c *Money) GetMoney() int {
	return c.val
}

func (c *Money) SetMoney(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Money) IsUpdated() bool {
	return c._updated
}

func (c *Money) resetUpdated() {
	c._updated = false
}

func (c *Money) GetColumnName() string {
	return "money"
}

func (c *Money) IsPrimaryKey() bool {
	return false
}

func (c *Money) GetValPointer() any {
	return &c.val
}

func (c *Money) GetTableType() coredb.TableType {
	return table
}

// NewAccountWithPK return new *Account with given PK
// PK column: "user_id","country_code"
func NewAccountWithPK(val PK) *Account {
	c := &Account{
		UserId{},
		Type{val: "free"},
		CountryCode{val: uint(0)},
		Money{val: int(0)},
	}
	c.UserId.val = val.UserId
	c.CountryCode.val = val.CountryCode
	return c
}

func (c *Account) Insert() error {
	sql := "INSERT IGNORE INTO `account` (`user_id`, `type`, `country_code`, `money`) values (?, ?, ?, ?)"

	result, err := coredb.Exec(sql, DBName, c.GetUserId(), c.GetType(), c.GetCountryCode(), c.GetMoney())

	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return coredb.ErrAvoidInsert
	}

	c.resetUpdated()
	return nil
}

func (c *Account) resetUpdated() {
	c.Type.resetUpdated()
	c.Money.resetUpdated()
}

func (obj *Account) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Type.IsUpdated() {
		updatedFields = append(updatedFields, "`type` = ?")
		params = append(params, obj.GetType())
	}
	if obj.Money.IsUpdated() {
		updatedFields = append(updatedFields, "`money` = ?")
		params = append(params, obj.GetMoney())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `account` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `country_code` = ?"
	params = append(params, obj.GetUserId(), obj.GetCountryCode())

	result, err := coredb.Exec(sql, DBName, params...)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affectedRows == 0 {
		return false, coredb.ErrAvoidUpdate
	}

	obj.resetUpdated()
	return true, nil
}

func Update(obj withPK) (bool, error) {
	var updatedFields []string
	var params []any
	var resetFuncs []func()

	val := reflect.ValueOf(obj).Elem()
	updatedFields = make([]string, 0, val.NumField())
	params = make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		col := val.Field(i).Addr().Interface()

		switch c := col.(type) {
		case *Type:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`type` = ?")
				params = append(params, c.GetType())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Money:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`money` = ?")
				params = append(params, c.GetMoney())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `account` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `country_code` = ?"
	params = append(params, obj.GetUserId(), obj.GetCountryCode())

	result, err := coredb.Exec(sql, DBName, params...)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if affectedRows == 0 {
		return false, coredb.ErrAvoidUpdate
	}

	for _, f := range resetFuncs {
		f()
	}
	return true, nil
}

func (obj *Account) Delete() error {
	sql := "DELETE FROM `account` WHERE `user_id` = ? and `country_code` = ?"

	_, err := coredb.Exec(sql, DBName, obj.GetUserId(), obj.GetCountryCode())
	return err
}

func Delete(obj withPK) error {
	sql := "DELETE FROM `account` WHERE `user_id` = ? and `country_code` = ?"

	_, err := coredb.Exec(sql, DBName, obj.GetUserId(), obj.GetCountryCode())
	return err
}
