package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

type Node interface {
	core.Node

	GetTag() string

	HasAttr(name string) bool
	GetAttr(name string) string
	SetAttr(name string, value string)
	DelAttr(name string)

	HasId() bool
	GetId() string
	SetId(id string)
	DelId()

	HasClass(name ...string) bool
	GetClass() []string
	AddClass(name ...string)
	DelClass(name ...string)

	GetStyles() maps.OMap[string, string]
	SetStyles(styles ...maps.KeyValue[string, string])

	HasStyle(name ...string) bool
	GetStyle(name string) string
	SetStyle(name string, value string)
	DelStyle(name ...string)

	GetById(id string) Node
	Select(query ...any) []Node
}
