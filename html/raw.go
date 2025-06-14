package html

import (
	"strings"

	"github.com/thegogod/mde/maps"
)

type Raw []byte

func (self Raw) GetSelector() string {
	return ""
}

func (self Raw) HasAttr(name string) bool {
	return false
}

func (self Raw) GetAttr(name string) string {
	return ""
}

func (self Raw) SetAttr(name string, value string) {
	return
}

func (self Raw) DelAttr(name string) {
	return
}

func (self Raw) HasId() bool {
	return false
}

func (self Raw) GetId() string {
	return ""
}

func (self Raw) SetId(id string) {
	return
}

func (self Raw) DelId() {
	return
}

func (self Raw) HasClass(name ...string) bool {
	return false
}

func (self Raw) GetClass() []string {
	return []string{}
}

func (self Raw) AddClass(name ...string) {
	return
}

func (self Raw) DelClass(name ...string) {
	return
}

func (self Raw) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self Raw) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self Raw) HasStyle(name ...string) bool {
	return false
}

func (self Raw) GetStyle(name string) string {
	return ""
}

func (self Raw) SetStyle(name string, value string) {
	return
}

func (self Raw) DelStyle(name ...string) {
	return
}

func (self Raw) String() string {
	return string(self)
}

func (self Raw) PrettyString(indent string) string {
	lines := strings.Split(string(self), "\n")
	return strings.Join(lines, "\n"+indent)
}

func (self Raw) Bytes() []byte {
	return self
}

func (self Raw) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self Raw) GetById(id string) Node {
	return nil
}

func (self Raw) Select(query ...any) []Node {
	return []Node{}
}
