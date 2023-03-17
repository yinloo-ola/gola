// Code generated by gola 0.1.1; DO NOT EDIT.

package song_user_favourites

import (
	"database/sql"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/olachat/gola/coredb"

	"time"
)

const DBName string = "testdata"
const TableName string = "song_user_favourites"

// SongUserFavourite represents `song_user_favourites` table
type SongUserFavourite struct {
	// User ID int unsigned
	UserId `json:"user_id"`
	// Song ID int unsigned
	SongId `json:"song_id"`
	// favourite remark varchar
	Remark `json:"remark"`
	// Is favourite tinyint(1)
	IsFavourite `json:"is_favourite"`
	// Create Time timestamp
	CreatedAt `json:"created_at"`
	// Last Update Time timestamp
	UpdatedAt `json:"updated_at"`
}
type PK struct {
	UserId uint
	SongId uint
}

type withPK interface {
	GetUserId() uint
	GetSongId() uint
}

// FetchByPK returns a row from `song_user_favourites` table with given primary key value
func FetchByPK(val PK) *SongUserFavourite {
	return coredb.FetchByPK[SongUserFavourite](DBName, TableName, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FetchFieldsByPK returns a row with selected fields from song_user_favourites table with given primary key value
func FetchFieldsByPK[T any](val PK) *T {
	return coredb.FetchByPK[T](DBName, TableName, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FindOne returns a row from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindOne(whereSQL string, params ...any) *SongUserFavourite {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[SongUserFavourite](DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFields[T any](whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOne[T](DBName, TableName, w)
}

// Find returns rows from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func Find(whereSQL string, params ...any) ([]*SongUserFavourite, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[SongUserFavourite](DBName, TableName, w)
}

// FindFields returns rows with selected fields from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindFields[T any](whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.Find[T](DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func Count(whereSQL string, params ...any) (int, error) {
	return coredb.QueryInt(DBName, "SELECT COUNT(*) FROM `song_user_favourites` "+whereSQL, params...)
}

// Column types

// UserId field
// User ID
type UserId struct {
	val uint
}

func (c *UserId) GetUserId() uint {
	return c.val
}

func (c *UserId) GetColumnName() string {
	return "user_id"
}

func (c *UserId) GetValPointer() any {
	return &c.val
}

func (c *UserId) getUserIdForDB() uint {
	return c.val
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

// SongId field
// Song ID
type SongId struct {
	val uint
}

func (c *SongId) GetSongId() uint {
	return c.val
}

func (c *SongId) GetColumnName() string {
	return "song_id"
}

func (c *SongId) GetValPointer() any {
	return &c.val
}

func (c *SongId) getSongIdForDB() uint {
	return c.val
}

func (c *SongId) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *SongId) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// Remark field
// favourite remark
type Remark struct {
	_updated bool
	val      string
}

func (c *Remark) GetRemark() string {
	return c.val
}

func (c *Remark) SetRemark(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Remark) IsUpdated() bool {
	return c._updated
}

func (c *Remark) resetUpdated() {
	c._updated = false
}

func (c *Remark) GetColumnName() string {
	return "remark"
}

func (c *Remark) GetValPointer() any {
	return &c.val
}

func (c *Remark) getRemarkForDB() string {
	return c.val
}

func (c *Remark) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *Remark) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// IsFavourite field
// Is favourite
type IsFavourite struct {
	_updated bool
	val      bool
}

func (c *IsFavourite) GetIsFavourite() bool {
	return c.val
}

func (c *IsFavourite) SetIsFavourite(val bool) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsFavourite) IsUpdated() bool {
	return c._updated
}

func (c *IsFavourite) resetUpdated() {
	c._updated = false
}

func (c *IsFavourite) GetColumnName() string {
	return "is_favourite"
}

func (c *IsFavourite) GetValPointer() any {
	return &c.val
}

func (c *IsFavourite) getIsFavouriteForDB() bool {
	return c.val
}

func (c *IsFavourite) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *IsFavourite) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// CreatedAt field
// Create Time
type CreatedAt struct {
	_updated bool
	val      time.Time
}

func (c *CreatedAt) GetCreatedAt() time.Time {
	return c.val
}

func (c *CreatedAt) SetCreatedAt(val time.Time) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *CreatedAt) IsUpdated() bool {
	return c._updated
}

func (c *CreatedAt) resetUpdated() {
	c._updated = false
}

func (c *CreatedAt) GetColumnName() string {
	return "created_at"
}

func (c *CreatedAt) GetValPointer() any {
	return &c.val
}

func (c *CreatedAt) getCreatedAtForDB() time.Time {
	return c.val
}

func (c *CreatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *CreatedAt) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// UpdatedAt field
// Last Update Time
type UpdatedAt struct {
	_updated bool
	val      time.Time
}

func (c *UpdatedAt) GetUpdatedAt() time.Time {
	return c.val
}

func (c *UpdatedAt) SetUpdatedAt(val time.Time) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UpdatedAt) IsUpdated() bool {
	return c._updated
}

func (c *UpdatedAt) resetUpdated() {
	c._updated = false
}

func (c *UpdatedAt) GetColumnName() string {
	return "updated_at"
}

func (c *UpdatedAt) GetValPointer() any {
	return &c.val
}

func (c *UpdatedAt) getUpdatedAtForDB() time.Time {
	return c.val
}

func (c *UpdatedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(&c.val)
}

func (c *UpdatedAt) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &c.val); err != nil {
		return err
	}

	return nil
}

// NewWithPK takes "user_id","song_id"
// and returns new *SongUserFavourite with given PK
func NewWithPK(val PK) *SongUserFavourite {
	c := &SongUserFavourite{
		UserId{},
		SongId{},
		Remark{},
		IsFavourite{val: true},
		CreatedAt{val: time.Now()},
		UpdatedAt{val: time.Now()},
	}
	c.UserId.val = val.UserId
	c.SongId.val = val.SongId
	return c
}

const insertWithoutPK string = "INSERT INTO `song_user_favourites` (`user_id`, `song_id`, `remark`, `is_favourite`, `created_at`, `updated_at`) values (?, ?, ?, ?, ?, ?)"

// Insert SongUserFavourite struct to `song_user_favourites` table
func (c *SongUserFavourite) Insert() error {
	var result sql.Result
	var err error
	result, err = coredb.Exec(DBName, insertWithoutPK, c.getUserIdForDB(), c.getSongIdForDB(), c.getRemarkForDB(), c.getIsFavouriteForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB())
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

func (c *SongUserFavourite) resetUpdated() {
	c.Remark.resetUpdated()
	c.IsFavourite.resetUpdated()
	c.CreatedAt.resetUpdated()
	c.UpdatedAt.resetUpdated()
}

// Update SongUserFavourite struct in `song_user_favourites` table
func (obj *SongUserFavourite) Update() (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Remark.IsUpdated() {
		updatedFields = append(updatedFields, "`remark` = ?")
		params = append(params, obj.GetRemark())
	}
	if obj.IsFavourite.IsUpdated() {
		updatedFields = append(updatedFields, "`is_favourite` = ?")
		params = append(params, obj.GetIsFavourite())
	}
	if obj.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`created_at` = ?")
		params = append(params, obj.GetCreatedAt())
	}
	if obj.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`updated_at` = ?")
		params = append(params, obj.GetUpdatedAt())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `song_user_favourites` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `song_id` = ?"
	params = append(params, obj.GetUserId(), obj.GetSongId())

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

// Update SongUserFavourite struct with given fields in `song_user_favourites` table
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
		case *Remark:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`remark` = ?")
				params = append(params, c.GetRemark())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsFavourite:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_favourite` = ?")
				params = append(params, c.GetIsFavourite())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CreatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`created_at` = ?")
				params = append(params, c.GetCreatedAt())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *UpdatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`updated_at` = ?")
				params = append(params, c.GetUpdatedAt())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `song_user_favourites` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `song_id` = ?"
	params = append(params, obj.GetUserId(), obj.GetSongId())

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

const deleteSql string = "DELETE FROM `song_user_favourites` WHERE `user_id` = ? and `song_id` = ?"

// DeleteByPK delete a row from song_user_favourites table with given primary key value
func DeleteByPK(val PK) error {
	_, err := coredb.Exec(DBName, deleteSql, val.UserId, val.SongId)
	return err
}
