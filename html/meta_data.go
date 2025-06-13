package html

import "github.com/thegogod/mde/maps"

type MetaData map[string]any

func (self MetaData) GetTag() string {
	return ":metadata"
}

func (self MetaData) HasAttr(name string) bool {
	return false
}

func (self MetaData) GetAttr(name string) string {
	return ""
}

func (self MetaData) SetAttr(name string, value string) {
	return
}

func (self MetaData) DelAttr(name string) {
	return
}

func (self MetaData) HasId() bool {
	return false
}

func (self MetaData) GetId() string {
	return ""
}

func (self MetaData) SetId(id string) {
	return
}

func (self MetaData) DelId() {
	return
}

func (self MetaData) HasClass(name ...string) bool {
	return false
}

func (self MetaData) GetClass() []string {
	return []string{}
}

func (self MetaData) AddClass(name ...string) {
	return
}

func (self MetaData) DelClass(name ...string) {
	return
}

func (self MetaData) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self MetaData) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self MetaData) HasStyle(name ...string) bool {
	return false
}

func (self MetaData) GetStyle(name string) string {
	return ""
}

func (self MetaData) SetStyle(name string, value string) {
	return
}

func (self MetaData) DelStyle(name ...string) {
	return
}

func (self MetaData) String() string {
	return ""
}

func (self MetaData) PrettyString(indent string) string {
	return ""
}

func (self MetaData) Bytes() []byte {
	return []byte{}
}

func (self MetaData) PrettyBytes(indent string) []byte {
	return []byte{}
}

func (self MetaData) GetById(id string) Node {
	return nil
}

func (self MetaData) GetByClass(classes ...string) []Node {
	return []Node{}
}

func (self MetaData) Select(query ...any) []Node {
	return []Node{}
}
