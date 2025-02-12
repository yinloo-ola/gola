// Code generated by gola 0.1.1; DO NOT EDIT.

package room

import (
	"fmt"
	"strings"

	"github.com/olachat/gola/coredb"
)

type orderBy int

type idxQuery[T any] struct {
	whereSql    string
	limitSql    string
	orders      []string
	whereParams []any
}

// order by enum & interface
const (
	IdAsc orderBy = iota
	IdDesc
	GroupAsc
	GroupDesc
	LangAsc
	LangDesc
	PriorityAsc
	PriorityDesc
	DeletedAsc
	DeletedDesc
)

func (q *idxQuery[T]) OrderBy(args ...orderBy) coredb.ReadQuery[T] {
	q.orders = make([]string, len(args))
	for i, arg := range args {
		switch arg {
		case IdAsc:
			q.orders[i] = "`id` asc"
		case IdDesc:
			q.orders[i] = "`id` desc"
		case GroupAsc:
			q.orders[i] = "`group` asc"
		case GroupDesc:
			q.orders[i] = "`group` desc"
		case LangAsc:
			q.orders[i] = "`lang` asc"
		case LangDesc:
			q.orders[i] = "`lang` desc"
		case PriorityAsc:
			q.orders[i] = "`priority` asc"
		case PriorityDesc:
			q.orders[i] = "`priority` desc"
		case DeletedAsc:
			q.orders[i] = "`deleted` asc"
		case DeletedDesc:
			q.orders[i] = "`deleted` desc"
		}
	}
	return q
}

func (q *idxQuery[T]) All() []*T {
	result, _ := coredb.Find[T](DBName, TableName, q)
	return result
}

func (q *idxQuery[T]) Limit(offset, limit int) []*T {
	q.limitSql = fmt.Sprintf(" limit %d, %d", offset, limit)
	result, _ := coredb.Find[T](DBName, TableName, q)
	return result
}

type order[T any] interface {
	OrderBy(args ...orderBy) coredb.ReadQuery[T]
}

type orderReadQuery[T any] interface {
	order[T]
	coredb.ReadQuery[T]
}

type iQuery[T any] interface {
	WhereGroupEQ(val uint) orderReadQuery[T]
	WhereGroupIN(vals ...uint) orderReadQuery[T]
	WhereLangEQ(val string) iQuery1[T]
	WhereLangIN(vals ...string) iQuery1[T]
	WherePriorityEQ(val float64) orderReadQuery[T]
	WherePriorityIN(vals ...float64) orderReadQuery[T]
	orderReadQuery[T]
}
type iQuery1[T any] interface {
	AndDeletedEQ(val bool) orderReadQuery[T]
	AndDeletedIN(vals ...bool) orderReadQuery[T]
	orderReadQuery[T]
}

type idxQuery1[T any] struct {
	*idxQuery[T]
}

func (q *idxQuery1[T]) AndDeletedEQ(val bool) orderReadQuery[T] {
	q.whereSql += " and `deleted` = ?"
	q.whereParams = append(q.whereParams, val)
	return q.idxQuery
}

func (q *idxQuery1[T]) AndDeletedIN(vals ...bool) orderReadQuery[T] {
	q.whereSql += " and `deleted` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q.idxQuery
}

// Find methods

// Select returns rows from `room` table with index awared query
func Select() iQuery[Room] {
	return new(idxQuery[Room])
}

// SelectFields returns rows with selected fields from `room` table with index awared query
func SelectFields[T any]() iQuery[T] {
	return new(idxQuery[T])
}

func (q *idxQuery[T]) WhereGroupEQ(val uint) orderReadQuery[T] {
	q.whereSql += " where `group` = ?"
	q.whereParams = append(q.whereParams, val)
	return q
}

func (q *idxQuery[T]) WhereGroupIN(vals ...uint) orderReadQuery[T] {
	q.whereSql = " where `group` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q
}

func (q *idxQuery[T]) WhereLangEQ(val string) iQuery1[T] {
	q.whereSql += " where `lang` = ?"
	q.whereParams = append(q.whereParams, val)
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) WhereLangIN(vals ...string) iQuery1[T] {
	q.whereSql = " where `lang` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return &idxQuery1[T]{q}
}

func (q *idxQuery[T]) WherePriorityEQ(val float64) orderReadQuery[T] {
	q.whereSql += " where `priority` = ?"
	q.whereParams = append(q.whereParams, val)
	return q
}

func (q *idxQuery[T]) WherePriorityIN(vals ...float64) orderReadQuery[T] {
	q.whereSql = " where `priority` in (" + coredb.GetParamPlaceHolder(len(vals)) + ")"
	for _, val := range vals {
		q.whereParams = append(q.whereParams, val)
	}
	return q
}

func (q *idxQuery[T]) GetWhere() (whereSql string, params []any) {
	var orderSql string
	if len(q.orders) > 0 {
		orderSql = " order by " + strings.Join(q.orders, ",")
	}
	return q.whereSql + orderSql + q.limitSql, q.whereParams
}
