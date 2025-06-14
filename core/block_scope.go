package core

import (
	"fmt"

	"github.com/thegogod/mde/reflect"
)

type BlockScope struct {
	parent Scope
	values map[string]reflect.Value
}

func NewBlockScope() *BlockScope {
	return &BlockScope{values: map[string]reflect.Value{}}
}

func (self *BlockScope) Create() Scope {
	child := NewBlockScope()
	child.parent = self
	return child
}

func (self BlockScope) HasLocal(key string) bool {
	_, exists := self.values[key]
	return exists
}

func (self BlockScope) Has(key string) bool {
	if self.HasLocal(key) {
		return true
	}

	if self.parent != nil {
		return self.parent.Has(key)
	}

	return false
}

func (self BlockScope) GetLocal(key string) (reflect.Value, error) {
	if !self.HasLocal(key) {
		return reflect.Null(), fmt.Errorf("identifier '%s' not found", key)
	}

	return self.values[key], nil
}

func (self BlockScope) Get(key string) (reflect.Value, error) {
	if !self.Has(key) {
		return reflect.Null(), fmt.Errorf("identifier '%s' not found", key)
	}

	if self.HasLocal(key) {
		return self.GetLocal(key)
	}

	return self.parent.Get(key)
}

func (self *BlockScope) Add(key string, value reflect.Value) error {
	if self.HasLocal(key) {
		return fmt.Errorf("cannot re-define existing identifier '%s'", key)
	}

	self.values[key] = value
	return nil
}

func (self *BlockScope) Set(key string, value reflect.Value) error {
	existing, err := self.Get(key)

	if err != nil {
		return err
	}

	if !existing.Type().Assignable(value.Type()) {
		return fmt.Errorf(
			"cannot assign value of type '%s' to type '%s'",
			value.Type().Name(),
			existing.Type().Name(),
		)
	}

	if self.HasLocal(key) {
		self.values[key] = value
		return nil
	}

	return self.parent.Set(key, value)
}

func (self *BlockScope) Del(key string) error {
	if !self.Has(key) {
		return fmt.Errorf("identifier '%s' not found", key)
	}

	if self.HasLocal(key) {
		delete(self.values, key)
		return nil
	}

	return self.parent.Del(key)
}
