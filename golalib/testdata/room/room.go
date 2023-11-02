// Code generated by gola 0.1.1; DO NOT EDIT.

package room

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
)

const DBName string = "testdata"
const TableName string = "room"

// Room represents `room` table
type Room struct {
	//  int(11) unsigned
	Id `json:"id"`
	//  int(11) unsigned
	Group `json:"group"`
	//  varchar(5)
	Lang `json:"lang"`
	//  double
	Priority `json:"priority"`
	//  tinyint(1) unsigned
	Deleted `json:"deleted"`
}

type withPK interface {
	GetId() uint
}

// FetchByPK returns a row from `room` table with given primary key value
func FetchByPK(val uint) *Room {
	return coredb.FetchByPK[Room](DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from room table with given primary key value
func FetchFieldsByPK[T any](val uint) *T {
	return coredb.FetchByPK[T](DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `room` table with given primary key values
func FetchByPKs(vals ...uint) []*Room {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[Room](DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `room` table with given primary key values
func FetchFieldsByPKs[T any](vals ...uint) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKs[T](DBName, TableName, "id", pks)
}

// FindOne returns a row from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOne(whereSQL string, params ...any) *Room {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[Room](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFields[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](DBName, TableName, w)
}

// Find returns rows from `room` table with arbitary where query
// whereSQL must start with "where ..."
func Find(whereSQL string, params ...any) ([]*Room, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[Room](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindFields[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt(DBName, "SELECT COUNT(*) FROM `room` "+whereSQL, params...)
}

// FetchByPK returns a row from `room` table with given primary key value
func FetchByPKFromMaster(val uint) *Room {
	return coredb.FetchByPKFromMaster[Room](DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from room table with given primary key value
func FetchFieldsByPKFromMaster[T any](val uint) *T {
	return coredb.FetchByPKFromMaster[T](DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `room` table with given primary key values
func FetchByPKsFromMaster(vals ...uint) []*Room {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMaster[Room](DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `room` table with given primary key values
func FetchFieldsByPKsFromMaster[T any](vals ...uint) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMaster[T](DBName, TableName, "id", pks)
}

// FindOne returns a row from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFromMaster(whereSQL string, params ...any) *Room {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMaster[Room](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsFromMaster[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMaster[T](DBName, TableName, w)
}

// Find returns rows from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindFromMaster(whereSQL string, params ...any) ([]*Room, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMaster[Room](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `room` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsFromMaster[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMaster[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountFromMaster(whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntFromMaster(DBName, "SELECT COUNT(*) FROM `room` "+whereSQL, params...)
}

// Column types

// Id field
type Id struct {
	isAssigned bool
	val        uint
}

func (c *Id) GetId() uint {
	return c.val
}

func (c *Id) GetColumnName() string {
	return "id"
}

func (c *Id) GetValPointer() any {
	return &c.val
}

func (c *Id) getIdForDB() uint {
	return c.val
}

func (c *Id) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Id) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Group field
type Group struct {
	_updated bool
	val      uint
}

func (c *Group) GetGroup() uint {
	return c.val
}

func (c *Group) SetGroup(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Group) IsUpdated() bool {
	return c._updated
}

func (c *Group) resetUpdated() {
	c._updated = false
}

func (c *Group) GetColumnName() string {
	return "group"
}

func (c *Group) GetValPointer() any {
	return &c.val
}

func (c *Group) getGroupForDB() uint {
	return c.val
}

func (c *Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Group) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Lang field
type Lang struct {
	_updated bool
	val      string
}

func (c *Lang) GetLang() string {
	return c.val
}

func (c *Lang) SetLang(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Lang) IsUpdated() bool {
	return c._updated
}

func (c *Lang) resetUpdated() {
	c._updated = false
}

func (c *Lang) GetColumnName() string {
	return "lang"
}

func (c *Lang) GetValPointer() any {
	return &c.val
}

func (c *Lang) getLangForDB() string {
	return c.val
}

func (c *Lang) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Lang) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Priority field
type Priority struct {
	_updated bool
	val      float64
}

func (c *Priority) GetPriority() float64 {
	return c.val
}

func (c *Priority) SetPriority(val float64) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Priority) IsUpdated() bool {
	return c._updated
}

func (c *Priority) resetUpdated() {
	c._updated = false
}

func (c *Priority) GetColumnName() string {
	return "priority"
}

func (c *Priority) GetValPointer() any {
	return &c.val
}

func (c *Priority) getPriorityForDB() float64 {
	return c.val
}

func (c *Priority) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Priority) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Deleted field
type Deleted struct {
	_updated bool
	val      bool
}

func (c *Deleted) GetDeleted() bool {
	return c.val
}

func (c *Deleted) SetDeleted(val bool) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Deleted) IsUpdated() bool {
	return c._updated
}

func (c *Deleted) resetUpdated() {
	c._updated = false
}

func (c *Deleted) GetColumnName() string {
	return "deleted"
}

func (c *Deleted) GetValPointer() any {
	return &c.val
}

func (c *Deleted) getDeletedForDB() bool {
	return c.val
}

func (c *Deleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Deleted) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// New return new *Room with default values
func New() *Room {
	return &Room{
		Id{},
		Group{val: uint(0)},
		Lang{val: "en"},
		Priority{val: float64(0)},
		Deleted{val: false},
	}
}

// NewWithPK takes "id"
// and returns new *Room with given PK
func NewWithPK(val uint) *Room {
	c := &Room{
		Id{},
		Group{val: uint(0)},
		Lang{val: "en"},
		Priority{val: float64(0)},
		Deleted{val: false},
	}
	c.Id.val = val
	c.Id.isAssigned = true
	return c
}

const insertWithoutPK string = "INSERT INTO `room` (`group`, `lang`, `priority`, `deleted`) values (?, ?, ?, ?)"
const insertWithPK string = "INSERT INTO `room` (`id`, `group`, `lang`, `priority`, `deleted`) values (?, ?, ?, ?, ?)"

// Insert Room struct to `room` table
func (c *Room) Insert() error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.Exec(DBName, insertWithPK, c.getIdForDB(), c.getGroupForDB(), c.getLangForDB(), c.getPriorityForDB(), c.getDeletedForDB())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.Exec(DBName, insertWithoutPK, c.getGroupForDB(), c.getLangForDB(), c.getPriorityForDB(), c.getDeletedForDB())
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		c.Id.val = uint(id)
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

func (c *Room) resetUpdated() {
	c.Group.resetUpdated()
	c.Lang.resetUpdated()
	c.Priority.resetUpdated()
	c.Deleted.resetUpdated()
}

// Update Room struct in `room` table
func (obj *Room) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Group.IsUpdated() {
		updatedFields = append(updatedFields, "`group` = ?")
		params = append(params, obj.getGroupForDB())
	}
	if obj.Lang.IsUpdated() {
		updatedFields = append(updatedFields, "`lang` = ?")
		params = append(params, obj.getLangForDB())
	}
	if obj.Priority.IsUpdated() {
		updatedFields = append(updatedFields, "`priority` = ?")
		params = append(params, obj.getPriorityForDB())
	}
	if obj.Deleted.IsUpdated() {
		updatedFields = append(updatedFields, "`deleted` = ?")
		params = append(params, obj.getDeletedForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `room` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

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

// Update Room struct with given fields in `room` table
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
		case *Group:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`group` = ?")
				params = append(params, c.getGroupForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Lang:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`lang` = ?")
				params = append(params, c.getLangForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Priority:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`priority` = ?")
				params = append(params, c.getPriorityForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Deleted:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`deleted` = ?")
				params = append(params, c.getDeletedForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `room` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

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

const deleteSql string = "DELETE FROM `room` WHERE `id` = ?"

// DeleteByPK delete a row from room table with given primary key value
func DeleteByPK(val uint) error {
	_, err := coredb.Exec(DBName, deleteSql, val)
	return err
}
