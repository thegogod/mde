package syntax

import (
	"fmt"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

const infoIcon = `<svg viewBox="0 0 14 16"><path fill-rule="evenodd" d="M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"></path></svg>`
const tipIcon = `<svg viewBox="0 0 12 16"><path fill-rule="evenodd" d="M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"></path></svg>`
const warnIcon = `<svg viewBox="0 0 16 16"><path fill-rule="evenodd" d="M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"></path></svg>`
const dangerIcon = `<svg viewBox="0 0 12 16"><path fill-rule="evenodd" d="M5.05.31c.81 2.17.41 3.38-.52 4.31C3.55 5.67 1.98 6.45.9 7.98c-1.45 2.05-1.7 6.53 3.53 7.7-2.2-1.16-2.67-4.52-.3-6.61-.61 2.03.53 3.33 1.94 2.86 1.39-.47 2.3.53 2.27 1.67-.02.78-.31 1.44-1.13 1.81 3.42-.59 4.78-3.42 4.78-5.56 0-2.84-2.53-3.22-1.25-5.61-1.52.13-2.03 1.13-1.89 2.75.09 1.08-1.02 1.8-1.86 1.33-.67-.41-.66-1.19-.06-1.78C8.18 5.31 8.68 2.45 5.05.32L5.03.3l.02.01z"></path></svg>`
const styles = `.alert {
	display: block;
	padding: 1rem 1rem;
	border-radius: 0.4rem;

	&.note {
		color: #474748;
		background-color: #fdfdfe;
		border: 0px solid #d4d5d8;
		border-left-width: 5px;
		box-shadow: 0 1px 2px 0 #0000001a;

		code {
			background-color: #ebedf026;
		}
	}

	&.info {
		color: #193c47;
		background-color: #eef9fd;
		border: 0px solid #4cb3d4;
		border-left-width: 5px;
		box-shadow: 0 1px 2px 0 #0000001a;

		code {
			background-color: #54c7ec26
		}
	}

	.header {
		display: flex;
		margin-bottom: .3rem;

		svg {
			width: 1.6em;
			height: 1.6em;
			fill: #193c47;
			display: inline-block;
			vertical-align: middle;
			margin-right: 0.4em;
			display: inline-block;
		}

		.title {
			margin: auto 0;
			font-weight: bold;
			text-transform: uppercase;
		}
	}

	.body {
	}
}`

type Alert struct{}

func (self Alert) IsBlock() bool {
	return true
}

func (self Alert) IsInline() bool {
	return true
}

func (self Alert) Name() string {
	return "alert"
}

func (self Alert) Select(parser core.Parser, iter *core.Iterator) bool {
	return iter.Match(core.LessThan)
}

func (self Alert) Parse(parser core.Parser, iter *core.Iterator) (core.Node, error) {
	node, err := parser.ParseSyntax("markdown", "html", parser, iter)

	if node == nil || err != nil {
		return node, err
	}

	el := node.(html.ParentNode)

	if el.GetSelector() != "alert" {
		return nil, iter.Curr().Error(fmt.Sprintf("expected 'alert', received '%s'", el.GetSelector()))
	}

	kind := "note"

	if el.HasAttr("type") {
		kind = el.GetAttr("type")
	}

	style := html.Style(styles)
	header := html.Div().WithClass("header")
	body := html.Div().WithClass("body")
	host := html.Div().WithClass("alert", kind).Push(header, body)

	switch kind {
	case "note":
		fallthrough
	case "info":
		header.Push(infoIcon)
		break
	case "tip":
		header.Push(tipIcon)
		break
	case "warn":
		header.Push(warnIcon)
		break
	case "danger":
		header.Push(dangerIcon)
		break
	default:
		return el, iter.Curr().Error(fmt.Sprintf("invalid alert type '%s'", kind))
	}

	header.Push(html.Div(kind).WithClass("title"))

	for _, child := range el.Children() {
		body.Push(child)
	}

	template := html.Template(style, host).WithAttr("shadowrootmode", "open")
	return template, nil
}
