package reflect_test

import (
	"testing"

	"github.com/thegogod/mde/reflect"
)

func TestUnion(t *testing.T) {
	t.Run("should have name", func(t *testing.T) {
		union := reflect.Union(reflect.NumberType{}, reflect.StringType{})

		if union.Name() != "number | string" {
			t.FailNow()
		}
	})

	t.Run("should be assignable", func(t *testing.T) {
		union := reflect.Union(reflect.NumberType{}, reflect.StringType{})

		if !union.Assignable(reflect.NumberType{}) {
			t.FailNow()
		}

		if !union.Assignable(reflect.StringType{}) {
			t.FailNow()
		}

		if union.Assignable(reflect.BoolType{}) {
			t.FailNow()
		}
	})
}
