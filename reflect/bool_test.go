package reflect_test

import (
	"testing"

	"github.com/thegogod/mde/reflect"
)

func TestBool(t *testing.T) {
	t.Run("should be equal", func(t *testing.T) {
		a := reflect.Bool(true)
		b := reflect.Bool(true)

		if !a.Equals(b) {
			t.Fatalf("expected a == b")
		}
	})

	t.Run("should not be equal", func(t *testing.T) {
		a := reflect.Bool(true)
		b := reflect.Bool(false)

		if a.Equals(b) {
			t.Fatalf("expected a != b")
		}
	})

	t.Run("string", func(t *testing.T) {
		t.Run("should be true", func(t *testing.T) {
			if reflect.Bool(true).String() != "true" {
				t.FailNow()
			}
		})

		t.Run("should be false", func(t *testing.T) {
			if reflect.Bool(false).String() != "false" {
				t.FailNow()
			}
		})
	})
}
