package html

import (
	"fmt"
	"strings"

	"github.com/thegogod/mde/maps"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Guides/Comments
type Comment []byte

func (self Comment) GetSelector() string {
	return ""
}

func (self Comment) HasAttr(name string) bool {
	return false
}

func (self Comment) GetAttr(name string) string {
	return ""
}

func (self Comment) SetAttr(name string, value string) {
	return
}

func (self Comment) DelAttr(name string) {
	return
}

func (self Comment) HasId() bool {
	return false
}

func (self Comment) GetId() string {
	return ""
}

func (self Comment) SetId(id string) {
	return
}

func (self Comment) DelId() {
	return
}

func (self Comment) HasClass(name ...string) bool {
	return false
}

func (self Comment) GetClass() []string {
	return []string{}
}

func (self Comment) AddClass(name ...string) {
	return
}

func (self Comment) DelClass(name ...string) {
	return
}

func (self Comment) GetStyles() maps.OMap[string, string] {
	return maps.OMap[string, string]{}
}

func (self Comment) SetStyles(styles ...maps.KeyValue[string, string]) {
	return
}

func (self Comment) HasStyle(name ...string) bool {
	return false
}

func (self Comment) GetStyle(name string) string {
	return ""
}

func (self Comment) SetStyle(name string, value string) {
	return
}

func (self Comment) DelStyle(name ...string) {
	return
}

func (self Comment) String() string {
	return fmt.Sprintf("<!-- %s -->", string(self))
}

func (self Comment) PrettyString(indent string) string {
	lines := strings.Split(string(self), "\n")
	formatted := strings.Join(lines, "\n"+indent)
	return fmt.Sprintf("<!--\n %s \n-->", formatted)
}

func (self Comment) Bytes() []byte {
	return []byte(self.String())
}

func (self Comment) PrettyBytes(indent string) []byte {
	return []byte(self.PrettyString(indent))
}

func (self Comment) GetById(id string) Node {
	return nil
}

func (self Comment) Select(query ...any) []Node {
	return []Node{}
}
