package html

import (
	"fmt"
	"strings"

	"github.com/thegogod/mde/core"
)

// https://developer.mozilla.org/en-US/docs/Web/HTML/Guides/Comments
type Comment []byte

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

func (self Comment) GetById(id string) core.Node {
	return nil
}

func (self Comment) GetByClass(classes ...string) []core.Node {
	return []core.Node{}
}
