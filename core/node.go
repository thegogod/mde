package core

import "reflect"

type Node interface {
	Render() (reflect.Value, error)
}
