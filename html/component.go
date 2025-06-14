package html

import "github.com/thegogod/mde/core"

type Component struct {
	Selector string
	Template string
	Host     map[string]string
	Scripts  []string
	Styles   []string
}

func (self Component) Build(parser core.Parser) (Node, error) {
	template := Template().WithAttr("shadowrootmode", "open")

	if self.Styles != nil {
		for _, style := range self.Styles {
			template.Push(Style(style))
		}
	}

	if self.Scripts != nil {
		for _, script := range self.Scripts {
			template.Push(Script(script))
		}
	}

	host := Elem(self.Selector)

	if self.Host != nil {
		for key, value := range self.Host {
			host.SetAttr(key, value)
		}
	}

	content, err := parser.Parse([]byte(self.Template))

	if content == nil || err != nil {
		return template, err
	}

	template.Push(content)
	return template, nil
}
