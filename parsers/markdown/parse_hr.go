package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

func (self *Parser) parseHr() (core.Node, error) {
	return html.Hr(), nil
}
