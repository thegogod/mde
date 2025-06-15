package html

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/maps"
)

type Host map[string]any

func (self Host) GetSelector() string {
	return ":host"
}

func (self Host) HasAttr(name string) bool {
	return false
}

func (self Host) GetAttr(name string) string {
	return ""
}

func (self Host) SetAttr(name string, value string) {
	return
}

func (self Host) DelAttr(name string) {
	return
}

func (self Host) HasId() bool {
	return false
}

func (self Host) GetId() string {
	return ""
}

func (self Host) SetId(id string) {
	return
}

func (self Host) DelId() {
	return
}

func (self Host) HasClass(name ...string) bool {
	return false
}

func (self Host) GetClass() []string {
	return []string{}
}

func (self Host) AddClass(name ...string) {
	return
}

func (self Host) DelClass(name ...string) {
	return
}

func (self Host) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self Host) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self Host) HasStyle(name ...string) bool {
	return false
}

func (self Host) GetStyle(name string) string {
	return ""
}

func (self Host) SetStyle(name string, value string) {
	return
}

func (self Host) DelStyle(name ...string) {
	return
}

func (self Host) Render(scope core.Scope) []byte {
	return []byte{}
}

func (self Host) RenderPretty(scope core.Scope, indent string) []byte {
	return []byte{}
}

func (self Host) GetById(id string) Node {
	return nil
}

func (self Host) Select(query ...any) []Node {
	stmt := Select()

	for _, q := range query {
		switch v := q.(type) {
		case SelectStatement:
			stmt.And(v)
			break
		case string:
			break
		default:
			panic("invalid selector type")
		}
	}

	nodes := []Node{}

	if stmt.Eval(self) {
		nodes = append(nodes, self)
	}

	return nodes
}
