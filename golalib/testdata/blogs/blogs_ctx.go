// Code generated by gola 0.1.1; DO NOT EDIT.

package blogs

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
)

// FetchByPK returns a row from `blogs` table with given primary key value
func FetchByPKCtx(ctx context.Context, val int) *Blog {
	return coredb.FetchByPKCtx[Blog](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from blogs table with given primary key value
func FetchFieldsByPKCtx[T any](ctx context.Context, val int) *T {
	return coredb.FetchByPKCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `blogs` table with given primary key values
func FetchByPKsCtx(ctx context.Context, vals ...int) []*Blog {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[Blog](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `blogs` table with given primary key values
func FetchFieldsByPKsCtx[T any](ctx context.Context, vals ...int) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneCtx(ctx context.Context, whereSQL string, params ...any) *Blog {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[Blog](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindCtx(ctx context.Context, whereSQL string, params ...any) ([]*Blog, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[Blog](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntCtx(ctx, DBName, "SELECT COUNT(*) FROM `blogs` "+whereSQL, params...)
}

// FetchByPK returns a row from `blogs` table with given primary key value
func FetchByPKFromMasterCtx(ctx context.Context, val int) *Blog {
	return coredb.FetchByPKFromMasterCtx[Blog](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchFieldsByPK returns a row with selected fields from blogs table with given primary key value
func FetchFieldsByPKFromMasterCtx[T any](ctx context.Context, val int) *T {
	return coredb.FetchByPKFromMasterCtx[T](ctx, DBName, TableName, []string{"id"}, val)
}

// FetchByPKs returns rows with from `blogs` table with given primary key values
func FetchByPKsFromMasterCtx(ctx context.Context, vals ...int) []*Blog {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[Blog](ctx, DBName, TableName, "id", pks)
}

// FetchFieldsByPKs returns rows with selected fields from `blogs` table with given primary key values
func FetchFieldsByPKsFromMasterCtx[T any](ctx context.Context, vals ...int) []*T {
	pks := coredb.GetAnySlice(vals)
	return coredb.FetchByPKsFromMasterCtx[T](ctx, DBName, TableName, "id", pks)
}

// FindOne returns a row from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFromMasterCtx(ctx context.Context, whereSQL string, params ...any) *Blog {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[Blog](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) *T {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindFromMasterCtx(ctx context.Context, whereSQL string, params ...any) ([]*Blog, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[Blog](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `blogs` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntFromMasterCtx(ctx, DBName, "SELECT COUNT(*) FROM `blogs` "+whereSQL, params...)
}

// Insert Blog struct to `blogs` table
func (c *Blog) InsertCtx(ctx context.Context) error {
	var result sql.Result
	var err error
	if c.Id.isAssigned {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithPK, c.getIdForDB(), c.getUserIdForDB(), c.getSlugForDB(), c.getTitleForDB(), c.getCategoryIdForDB(), c.getIsPinnedForDB(), c.getIsVipForDB(), c.getCountryForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB())
		if err != nil {
			return err
		}
	} else {
		result, err = coredb.ExecCtx(ctx, DBName, insertWithoutPK, c.getUserIdForDB(), c.getSlugForDB(), c.getTitleForDB(), c.getCategoryIdForDB(), c.getIsPinnedForDB(), c.getIsVipForDB(), c.getCountryForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB())
		if err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		c.Id.val = int(id)
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

// Update Blog struct in `blogs` table
func (obj *Blog) UpdateCtx(ctx context.Context) (bool, error) {
	var updatedFields []string
	var params []any
	if obj.UserId.IsUpdated() {
		updatedFields = append(updatedFields, "`user_id` = ?")
		params = append(params, obj.getUserIdForDB())
	}
	if obj.Slug.IsUpdated() {
		updatedFields = append(updatedFields, "`slug` = ?")
		params = append(params, obj.getSlugForDB())
	}
	if obj.Title.IsUpdated() {
		updatedFields = append(updatedFields, "`title` = ?")
		params = append(params, obj.getTitleForDB())
	}
	if obj.CategoryId.IsUpdated() {
		updatedFields = append(updatedFields, "`category_id` = ?")
		params = append(params, obj.getCategoryIdForDB())
	}
	if obj.IsPinned.IsUpdated() {
		updatedFields = append(updatedFields, "`is_pinned` = ?")
		params = append(params, obj.getIsPinnedForDB())
	}
	if obj.IsVip.IsUpdated() {
		updatedFields = append(updatedFields, "`is_vip` = ?")
		params = append(params, obj.getIsVipForDB())
	}
	if obj.Country.IsUpdated() {
		updatedFields = append(updatedFields, "`country` = ?")
		params = append(params, obj.getCountryForDB())
	}
	if obj.CreatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`created_at` = ?")
		params = append(params, obj.getCreatedAtForDB())
	}
	if obj.UpdatedAt.IsUpdated() {
		updatedFields = append(updatedFields, "`updated_at` = ?")
		params = append(params, obj.getUpdatedAtForDB())
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `blogs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.ExecCtx(ctx, DBName, sql, params...)
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

// Update Blog struct with given fields in `blogs` table
func UpdateCtx(ctx context.Context, obj withPK) (bool, error) {
	var updatedFields []string
	var params []any
	var resetFuncs []func()

	val := reflect.ValueOf(obj).Elem()
	updatedFields = make([]string, 0, val.NumField())
	params = make([]any, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		col := val.Field(i).Addr().Interface()

		switch c := col.(type) {
		case *UserId:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`user_id` = ?")
				params = append(params, c.getUserIdForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Slug:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`slug` = ?")
				params = append(params, c.getSlugForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Title:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`title` = ?")
				params = append(params, c.getTitleForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CategoryId:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`category_id` = ?")
				params = append(params, c.getCategoryIdForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsPinned:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_pinned` = ?")
				params = append(params, c.getIsPinnedForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsVip:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_vip` = ?")
				params = append(params, c.getIsVipForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *Country:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`country` = ?")
				params = append(params, c.getCountryForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *CreatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`created_at` = ?")
				params = append(params, c.getCreatedAtForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *UpdatedAt:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`updated_at` = ?")
				params = append(params, c.getUpdatedAtForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		}
	}

	if len(updatedFields) == 0 {
		return false, nil
	}

	sql := "UPDATE `blogs` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `id` = ?"
	params = append(params, obj.GetId())

	result, err := coredb.ExecCtx(ctx, DBName, sql, params...)
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

// DeleteByPK delete a row from blogs table with given primary key value
func DeleteByPKCtx(ctx context.Context, val int) error {
	_, err := coredb.ExecCtx(ctx, DBName, deleteSql, val)
	return err
}
