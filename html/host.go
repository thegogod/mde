package html

import (
	"github.com/thegogod/mde/core"
)

type Host map[string]any

func (self Host) String() string {
	return ""
}

func (self Host) PrettyString(indent string) string {
	return ""
}

func (self Host) Bytes() []byte {
	return []byte{}
}

func (self Host) PrettyBytes(indent string) []byte {
	return []byte{}
}

func (self Host) GetById(id string) core.Node {
	return nil
}

func (self Host) GetByClass(classes ...string) []core.Node {
	return []core.Node{}
}
