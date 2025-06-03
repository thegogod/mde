package markdown

import (
	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

func (self *Parser) parseBr() (core.Node, error) {
	return html.Br(), nil
}
