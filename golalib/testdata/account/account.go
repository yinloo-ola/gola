// Code generated by gola 0.1.1; DO NOT EDIT.

package account

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"
)

const DBName string = "testdata"
const TableName string = "account"

// Account represents `account` table
type Account struct {
	//  int
	UserId `json:"user_id"`
	// user account type enum('free','vip')
	Type `json:"type"`
	// user country code mediumint unsigned
	CountryCode `json:"country_code"`
	// Account money int
	Money `json:"money"`
}
type PK struct {
	UserId      int
	CountryCode uint
}

type withPK interface {
	GetUserId() int
	GetCountryCode() uint
}

// FetchByPK returns a row from `account` table with given primary key value
func FetchByPK(val PK) *Account {
	return coredb.FetchByPK[Account](DBName, TableName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FetchFieldsByPK returns a row with selected fields from account table with given primary key value
func FetchFieldsByPK[T any](val PK) *T {
	return coredb.FetchByPK[T](DBName, TableName, []string{"user_id", "country_code"}, val.UserId, val.CountryCode)
}

// FindOne returns a row from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindOne(whereSQL string, params ...any) *Account {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[Account](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFields[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](DBName, TableName, w)
}

// Find returns rows from `account` table with arbitary where query
// whereSQL must start with "where ..."
func Find(whereSQL string, params ...any) ([]*Account, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[Account](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `account` table with arbitary where query
// whereSQL must start with "where ..."
func FindFields[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt(DBName, "SELECT COUNT(*) FROM `account` "+whereSQL, params...)
}

// Column types
type AccountType string

const (
	AccountTypeFree AccountType = "free"
	AccountTypeVip  AccountType = "vip"
)

// UserId field
type UserId struct {
	val int
}

func (c *UserId) GetUserId() int {
	return c.val
}

func (c *UserId) GetColumnName() string {
	return "user_id"
}

func (c *UserId) GetValPointer() any {
	return &c.val
}

func (c *UserId) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *UserId) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *Type) GetValPointer() any {
	return &c.val
}

func (c *Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Type) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *CountryCode) GetValPointer() any {
	return &c.val
}

func (c *CountryCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *CountryCode) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
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

func (c *Money) GetValPointer() any {
	return &c.val
}

func (c *Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Money) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// NewWithPK takes "user_id","country_code"
// and returns new *Account with given PK
func NewWithPK(val PK) *Account {
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

const insertWithoutPK string = "INSERT IGNORE INTO `account` (`user_id`, `type`, `country_code`, `money`) values (?, ?, ?, ?)"

// Insert Account struct to `account` table
func (c *Account) Insert() error {
	var result sql.Result
	var err error
	result, err = coredb.Exec(DBName, insertWithoutPK, c.GetUserId(), c.GetType(), c.GetCountryCode(), c.GetMoney())
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

// Update Account struct in `account` table
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

	result, err := coredb.Exec(DBName, sql, params...)
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

// Update Account struct with given fields in `account` table
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

	result, err := coredb.Exec(DBName, sql, params...)
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

const deleteSql string = "DELETE FROM `account` WHERE `user_id` = ? and `country_code` = ?"

// DeleteByPK delete a row from account table with given primary key value
func DeleteByPK(val PK) error {
	_, err := coredb.Exec(DBName, deleteSql, val.UserId, val.CountryCode)
	return err
}
