package query

import (
	"testing"
)

func TestGetPrimaryKey(t *testing.T) {
	GetPrimaryKey()
}

func TestQueryOne(t *testing.T) {
	QueryOne()
}

func TestQueryToMap(t *testing.T) {
	QueryToMap()
}

func TestQueryPluck(t *testing.T) {
	QueryPluck()
}

func TestQuerySelect(t *testing.T) {
	QuerySelect()
}

func TestQuerySelectEXP(t *testing.T) {
	QuerySelectEXP()
}

func TestQueryDistinct(t *testing.T) {
	QueryDistinct()
}

func TestWhereMethod(t *testing.T) {
	WhereMethod()
}

func TestWhereType(t *testing.T) {
	WhereType()
}

func TestPlaceHolder(t *testing.T) {
	PlaceHolder()
}

func TestOrderBy(t *testing.T) {
	OrderBy()
}

func TestPageOperation(t *testing.T) {
	PageOperation(Page{
		Page:     3,
		PageSize: 15,
	})
}

func TestLocking(t *testing.T) {
	Locking()
}

func TestSubQuery(t *testing.T) {
	SubQuery()
}

func TestQueryHook(t *testing.T) {
	QueryHook()
}
