package syntax

import "github.com/thegogod/mde/core"

var All []core.Syntax = []core.Syntax{
	H1{},
	H2{},
	H3{},
	H4{},
	H5{},
	H6{},
	HorizontalRule{},
	CodeBlock{},
	BlockQuote{},
	UnOrderedList{},
	OrderedList{},
	Table{},
	Paragraph{},

	Bold{},
	BoldAlt{},
	Italic{},
	ItalicAlt{},
	Strike{},
	StrikeAlt{},
	Highlight{},
	Code{},
	Emoji{},
	Link{},
	Url{},
	Image{},
	Break{},
	ListItem{},
	NewLine{},
}
