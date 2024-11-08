// Code generated by github.com/weightwave/gen. DO NOT EDIT.
// Code generated by github.com/weightwave/gen. DO NOT EDIT.
// Code generated by github.com/weightwave/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Input struct {
	Args []interface{}
}

type Expectation struct {
	Ret []interface{}
}

type TestCase struct {
	Input
	Expectation
}

const _gen_test_db_name = "gen_test.db"

var _gen_test_db *gorm.DB
var _gen_test_once sync.Once

func init() {
	InitializeDB()
	_gen_test_db.AutoMigrate(&_another{})
}

func InitializeDB() {
	_gen_test_once.Do(func() {
		var err error
		_gen_test_db, err = gorm.Open(sqlite.Open(_gen_test_db_name), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open sqlite %q fail: %w", _gen_test_db_name, err))
		}
	})
}

func assert(t *testing.T, methodName string, res, exp interface{}) {
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("%v() gotResult = %v, want %v", methodName, res, exp)
	}
}

type _another struct {
	ID uint64 `gorm:"primaryKey"`
}

func (*_another) TableName() string { return "another_for_unit_test" }

func Test_Available(t *testing.T) {
	if !Use(_gen_test_db).Available() {
		t.Errorf("query.Available() == false")
	}
}

func Test_WithContext(t *testing.T) {
	query := Use(_gen_test_db)
	if !query.Available() {
		t.Errorf("query Use(_gen_test_db) fail: query.Available() == false")
	}

	type Content string
	var key, value Content = "gen_tag", "unit_test"
	qCtx := query.WithContext(context.WithValue(context.Background(), key, value))

	for _, ctx := range []context.Context{
		qCtx.Bank.UnderlyingDB().Statement.Context,
		qCtx.CreditCard.UnderlyingDB().Statement.Context,
		qCtx.Customer.UnderlyingDB().Statement.Context,
		qCtx.Person.UnderlyingDB().Statement.Context,
		qCtx.User.UnderlyingDB().Statement.Context,
	} {
		if v := ctx.Value(key); v != value {
			t.Errorf("get value from context fail, expect %q, got %q", value, v)
		}
	}
}

func Test_Transaction(t *testing.T) {
	query := Use(_gen_test_db)
	if !query.Available() {
		t.Errorf("query Use(_gen_test_db) fail: query.Available() == false")
	}

	err := query.Transaction(func(tx *Query) error { return nil })
	if err != nil {
		t.Errorf("query.Transaction execute fail: %s", err)
	}

	tx := query.Begin()

	err = tx.SavePoint("point")
	if err != nil {
		t.Errorf("query tx SavePoint fail: %s", err)
	}
	err = tx.RollbackTo("point")
	if err != nil {
		t.Errorf("query tx RollbackTo fail: %s", err)
	}
	err = tx.Commit()
	if err != nil {
		t.Errorf("query tx Commit fail: %s", err)
	}

	err = query.Begin().Rollback()
	if err != nil {
		t.Errorf("query tx Rollback fail: %s", err)
	}
}
