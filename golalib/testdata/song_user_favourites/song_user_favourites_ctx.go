// Code generated by gola 0.1.1; DO NOT EDIT.

package song_user_favourites

import (
	"context"
	"database/sql"
	"reflect"
	"strings"

	"github.com/olachat/gola/v2/coredb"
)

// FetchByPK returns a row from `song_user_favourites` table with given primary key value
func FetchByPKCtx(ctx context.Context, val PK) (*SongUserFavourite, error) {
	return coredb.FetchByPKCtx[SongUserFavourite](ctx, DBName, TableName, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FetchFieldsByPK returns a row with selected fields from song_user_favourites table with given primary key value
func FetchFieldsByPKCtx[T any](ctx context.Context, val PK) (*T, error) {
	return coredb.FetchByPKCtx[T](ctx, DBName, TableName, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FindOne returns a row from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneCtx(ctx context.Context, whereSQL string, params ...any) (*SongUserFavourite, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[SongUserFavourite](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindCtx(ctx context.Context, whereSQL string, params ...any) ([]*SongUserFavourite, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[SongUserFavourite](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntCtx(ctx, DBName, "SELECT COUNT(*) FROM `song_user_favourites` "+whereSQL, params...)
}

// FetchByPK returns a row from `song_user_favourites` table with given primary key value
func FetchByPKFromMasterCtx(ctx context.Context, val PK) (*SongUserFavourite, error) {
	return coredb.FetchByPKFromMasterCtx[SongUserFavourite](ctx, DBName, TableName, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FetchFieldsByPK returns a row with selected fields from song_user_favourites table with given primary key value
func FetchFieldsByPKFromMasterCtx[T any](ctx context.Context, val PK) (*T, error) {
	return coredb.FetchByPKFromMasterCtx[T](ctx, DBName, TableName, []string{"user_id", "song_id"}, val.UserId, val.SongId)
}

// FindOne returns a row from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (*SongUserFavourite, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[SongUserFavourite](ctx, DBName, TableName, w)
}

// FindOneFields returns a row with selected fields from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindOneFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) (*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindOneFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Find returns rows from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindFromMasterCtx(ctx context.Context, whereSQL string, params ...any) ([]*SongUserFavourite, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[SongUserFavourite](ctx, DBName, TableName, w)
}

// FindFields returns rows with selected fields from `song_user_favourites` table with arbitary where query
// whereSQL must start with "where ..."
func FindFieldsFromMasterCtx[T any](ctx context.Context, whereSQL string, params ...any) ([]*T, error) {
	w := coredb.NewWhere(whereSQL, params...)
	return coredb.FindFromMasterCtx[T](ctx, DBName, TableName, w)
}

// Count returns select count(*) with arbitary where query
// whereSQL must start with "where ..."
func CountFromMasterCtx(ctx context.Context, whereSQL string, params ...any) (int, error) {
	return coredb.QueryIntFromMasterCtx(ctx, DBName, "SELECT COUNT(*) FROM `song_user_favourites` "+whereSQL, params...)
}

// Insert SongUserFavourite struct to `song_user_favourites` table
func (c *SongUserFavourite) InsertCtx(ctx context.Context) error {
	var result sql.Result
	var err error
	result, err = coredb.ExecCtx(ctx, DBName, insertWithoutPK, c.getUserIdForDB(), c.getSongIdForDB(), c.getRemarkForDB(), c.getIsFavouriteForDB(), c.getCreatedAtForDB(), c.getUpdatedAtForDB())
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

// Update SongUserFavourite struct in `song_user_favourites` table
func (obj *SongUserFavourite) UpdateCtx(ctx context.Context) (bool, error) {
	var updatedFields []string
	var params []any
	if obj.Remark.IsUpdated() {
		updatedFields = append(updatedFields, "`remark` = ?")
		params = append(params, obj.getRemarkForDB())
	}
	if obj.IsFavourite.IsUpdated() {
		updatedFields = append(updatedFields, "`is_favourite` = ?")
		params = append(params, obj.getIsFavouriteForDB())
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

	sql := "UPDATE `song_user_favourites` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `song_id` = ?"
	params = append(params, obj.GetUserId(), obj.GetSongId())

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

// Update SongUserFavourite struct with given fields in `song_user_favourites` table
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
		case *Remark:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`remark` = ?")
				params = append(params, c.getRemarkForDB())
				resetFuncs = append(resetFuncs, c.resetUpdated)
			}
		case *IsFavourite:
			if c.IsUpdated() {
				updatedFields = append(updatedFields, "`is_favourite` = ?")
				params = append(params, c.getIsFavouriteForDB())
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

	sql := "UPDATE `song_user_favourites` SET "
	sql = sql + strings.Join(updatedFields, ",") + " WHERE `user_id` = ? and `song_id` = ?"
	params = append(params, obj.GetUserId(), obj.GetSongId())

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

// DeleteByPK delete a row from song_user_favourites table with given primary key value
func DeleteByPKCtx(ctx context.Context, val PK) error {
	_, err := coredb.ExecCtx(ctx, DBName, deleteSql, val.UserId, val.SongId)
	return err
}