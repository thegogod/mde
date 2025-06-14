package core

import "github.com/thegogod/mde/reflect"

type ReadScope interface {
	HasLocal(key string) bool
	Has(key string) bool

	GetLocal(key string) (reflect.Value, error)
	Get(key string) (reflect.Value, error)
}

type WriteScope interface {
	Add(key string, value reflect.Value) error
	Set(key string, value reflect.Value) error
	Del(key string) error
}

type Scope interface {
	ReadScope
	WriteScope

	Create() Scope
}
