package mde

type Scope[T any] interface {
	Parent() Scope[T]

	Has(key string) bool
	HasLocal(key string) bool

	Get(key string) T
	GetLocal(key string) T

	Set(key string, value T) error
	Define(key string, value T) error
}
