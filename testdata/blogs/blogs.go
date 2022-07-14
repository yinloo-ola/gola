// Code generated by gola 0.0.2; DO NOT EDIT.

package blogs

import (
	"database/sql"
	"strings"

	"github.com/olachat/gola/corelib"
)

var _db *sql.DB

func Setup(db *sql.DB) {
	_db = db
}

// Blog represents blogs table
type Blog struct {
	//  int
	Id
	// User Id int
	UserId
	// Slug varchar(255)
	Slug
	// Title varchar(255)
	Title
	// Category Id int
	CategoryId
	// Is pinned to top tinyint
	IsPinned
	// Is VIP reader only tinyint
	IsVip
	// Country of the blog user varchar(255)
	Country
	// Created Timestamp int unsigned
	CreatedAt
	// Updated Timestamp int unsigned
	UpdatedAt
}

type BlogTable struct{}

func (*BlogTable) GetTableName() string {
	return "blogs"
}

var table *BlogTable

// FetchBlogByPKs returns a row from blogs table with given primary key value
func FetchBlogByPK(val int) *Blog {
	return corelib.FetchByPK[Blog](val, "id", _db)
}

// FetchByPKs returns a row with selected fields from blogs table with given primary key value
func FetchByPK[T any](val int) *T {
	return corelib.FetchByPK[T](val, "id", _db)
}

// FetchBlogByPKs returns rows with from blogs table with given primary key values
func FetchBlogByPKs(vals ...int) []*Blog {
	pks := corelib.GetAnySlice(vals)
	return corelib.FetchByPKs[Blog](pks, "id", _db)
}

// FetchByPKs returns rows with selected fields from blogs table with given primary key values
func FetchByPKs[T any](vals ...int) []*T {
	pks := corelib.GetAnySlice(vals)
	return corelib.FetchByPKs[T](pks, "id", _db)
}

// Column types

// Id field
//
type Id struct {
	_updated bool
	val      int
}

func (c *Id) GetId() int {
	return c.val
}

func (c *Id) SetId(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Id) GetColumnName() string {
	return "id"
}

func (c *Id) IsUpdated() bool {
	return c._updated
}

func (c *Id) IsPrimaryKey() bool {
	return true
}

func (c *Id) GetValPointer() interface{} {
	return &c.val
}

func (c *Id) GetTableType() corelib.TableType {
	return table
}

// UserId field
// User Id
type UserId struct {
	_updated bool
	val      int
}

func (c *UserId) GetUserId() int {
	return c.val
}

func (c *UserId) SetUserId(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UserId) GetColumnName() string {
	return "user_id"
}

func (c *UserId) IsUpdated() bool {
	return c._updated
}

func (c *UserId) IsPrimaryKey() bool {
	return false
}

func (c *UserId) GetValPointer() interface{} {
	return &c.val
}

func (c *UserId) GetTableType() corelib.TableType {
	return table
}

// Slug field
// Slug
type Slug struct {
	_updated bool
	val      string
}

func (c *Slug) GetSlug() string {
	return c.val
}

func (c *Slug) SetSlug(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Slug) GetColumnName() string {
	return "slug"
}

func (c *Slug) IsUpdated() bool {
	return c._updated
}

func (c *Slug) IsPrimaryKey() bool {
	return false
}

func (c *Slug) GetValPointer() interface{} {
	return &c.val
}

func (c *Slug) GetTableType() corelib.TableType {
	return table
}

// Title field
// Title
type Title struct {
	_updated bool
	val      string
}

func (c *Title) GetTitle() string {
	return c.val
}

func (c *Title) SetTitle(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Title) GetColumnName() string {
	return "title"
}

func (c *Title) IsUpdated() bool {
	return c._updated
}

func (c *Title) IsPrimaryKey() bool {
	return false
}

func (c *Title) GetValPointer() interface{} {
	return &c.val
}

func (c *Title) GetTableType() corelib.TableType {
	return table
}

// CategoryId field
// Category Id
type CategoryId struct {
	_updated bool
	val      int
}

func (c *CategoryId) GetCategoryId() int {
	return c.val
}

func (c *CategoryId) SetCategoryId(val int) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *CategoryId) GetColumnName() string {
	return "category_id"
}

func (c *CategoryId) IsUpdated() bool {
	return c._updated
}

func (c *CategoryId) IsPrimaryKey() bool {
	return false
}

func (c *CategoryId) GetValPointer() interface{} {
	return &c.val
}

func (c *CategoryId) GetTableType() corelib.TableType {
	return table
}

// IsPinned field
// Is pinned to top
type IsPinned struct {
	_updated bool
	val      int8
}

func (c *IsPinned) GetIsPinned() int8 {
	return c.val
}

func (c *IsPinned) SetIsPinned(val int8) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsPinned) GetColumnName() string {
	return "is_pinned"
}

func (c *IsPinned) IsUpdated() bool {
	return c._updated
}

func (c *IsPinned) IsPrimaryKey() bool {
	return false
}

func (c *IsPinned) GetValPointer() interface{} {
	return &c.val
}

func (c *IsPinned) GetTableType() corelib.TableType {
	return table
}

// IsVip field
// Is VIP reader only
type IsVip struct {
	_updated bool
	val      int8
}

func (c *IsVip) GetIsVip() int8 {
	return c.val
}

func (c *IsVip) SetIsVip(val int8) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *IsVip) GetColumnName() string {
	return "is_vip"
}

func (c *IsVip) IsUpdated() bool {
	return c._updated
}

func (c *IsVip) IsPrimaryKey() bool {
	return false
}

func (c *IsVip) GetValPointer() interface{} {
	return &c.val
}

func (c *IsVip) GetTableType() corelib.TableType {
	return table
}

// Country field
// Country of the blog user
type Country struct {
	_updated bool
	val      string
}

func (c *Country) GetCountry() string {
	return c.val
}

func (c *Country) SetCountry(val string) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *Country) GetColumnName() string {
	return "country"
}

func (c *Country) IsUpdated() bool {
	return c._updated
}

func (c *Country) IsPrimaryKey() bool {
	return false
}

func (c *Country) GetValPointer() interface{} {
	return &c.val
}

func (c *Country) GetTableType() corelib.TableType {
	return table
}

// CreatedAt field
// Created Timestamp
type CreatedAt struct {
	_updated bool
	val      uint
}

func (c *CreatedAt) GetCreatedAt() uint {
	return c.val
}

func (c *CreatedAt) SetCreatedAt(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *CreatedAt) GetColumnName() string {
	return "created_at"
}

func (c *CreatedAt) IsUpdated() bool {
	return c._updated
}

func (c *CreatedAt) IsPrimaryKey() bool {
	return false
}

func (c *CreatedAt) GetValPointer() interface{} {
	return &c.val
}

func (c *CreatedAt) GetTableType() corelib.TableType {
	return table
}

// UpdatedAt field
// Updated Timestamp
type UpdatedAt struct {
	_updated bool
	val      uint
}

func (c *UpdatedAt) GetUpdatedAt() uint {
	return c.val
}

func (c *UpdatedAt) SetUpdatedAt(val uint) bool {
	if c.val == val {
		return false
	}
	c._updated = true
	c.val = val
	return true
}

func (c *UpdatedAt) GetColumnName() string {
	return "updated_at"
}

func (c *UpdatedAt) IsUpdated() bool {
	return c._updated
}

func (c *UpdatedAt) IsPrimaryKey() bool {
	return false
}

func (c *UpdatedAt) GetValPointer() interface{} {
	return &c.val
}

func (c *UpdatedAt) GetTableType() corelib.TableType {
	return table
}

func NewBlog() *Blog {
	return &Blog{
		Id{},
		UserId{val: int(0)},
		Slug{val: ""},
		Title{val: ""},
		CategoryId{val: int(0)},
		IsPinned{val: int8(0)},
		IsVip{val: int8(0)},
		Country{val: ""},
		CreatedAt{val: uint(0)},
		UpdatedAt{val: uint(0)},
	}
}

func (c *Blog) Insert() error {
	sql := `INSERT INTO blogs (user_id, slug, title, category_id, is_pinned, is_vip, country, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := corelib.Exec(sql, _db, c.GetUserId(), c.GetSlug(), c.GetTitle(), c.GetCategoryId(), c.GetIsPinned(), c.GetIsVip(), c.GetCountry(), c.GetCreatedAt(), c.GetUpdatedAt())

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	c.SetId(int(id))
	return nil
}

func (c *Blog) Update() (bool, error) {
	var updatedFields []string
	var params []interface{}
	if c.UserId.IsUpdated() {
		updatedFields = append(updatedFields, "user_id = ?")
		params = append(params, c.GetUserId())
	}
	if c.Slug.IsUpdated() {
		updatedFields = append(updatedFields, "slug = ?")
		params = append(params, c.GetSlug())
	}
	if c.Title.IsUpdated() {
		updatedFields = append(updatedFields, "title = ?")
		params = append(params, c.GetTitle())
	}
	if c.CategoryId.IsUpdated() {
		updatedFields = append(updatedFields, "category_id = ?")
		params = append(params, c.GetCategoryId())
	}
	if c.IsPinned.IsUpdated() {
		updatedFields = append(updatedFields, "is_pinned = ?")
		params = append(params, c.GetIsPinned())
	}
	if c.IsVip.IsUpdated() {
		updatedFields = append(updatedFields, "is_vip = ?")
		params = append(params, c.GetIsVip())
	}
	if c.Country.IsUpdated() {
		updatedFields = append(updatedFields, "country = ?")
		params = append(params, c.GetCountry())
	}
	if c.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "created_at = ?")
		params = append(params, c.GetCreatedAt())
	}
	if c.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "updated_at = ?")
		params = append(params, c.GetUpdatedAt())
	}

	sql := `update blogs set `

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql = sql + strings.Join(updatedFields, ",") + " where id = ?"
	params = append(params, c.GetId())

	_, err := corelib.Exec(sql, _db, params...)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Blog) Delete() error {
	sql := `DELETE FROM blogs WHERE id = ?`

	_, err := corelib.Exec(sql, _db, c.GetId())
	return err
}

func Update[T any](obj *T) (bool, error) {
	return corelib.Update(obj, _db)
}
