package html_test

import (
	"testing"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/html"
)

func TestElement(t *testing.T) {
	t.Run("should render", func(t *testing.T) {
		el := html.Elem("div")
		el.Attr("id", "123")
		el.Styles(core.Pair("color", "red"), core.Pair("display", "block"))
		html := el.String()

		if html != `<div id="123" style="color: red;display: block;"></div>` {
			t.Error(html)
		}
	})

	t.Run("should render pretty", func(t *testing.T) {
		el := html.Elem("div")
		el.Attr("id", "123")
		el.Attr("style", "color: red;display: block;")
		html := el.PrettyString(" ")

		if html != "<div id=\"123\" style=\"color: red;display: block;\">\n</div>" {
			t.Error(html)
		}
	})

	t.Run("should render with children", func(t *testing.T) {
		el := html.Elem("div")
		el.Attr("id", "123")
		el.Attr("class", "main")
		el.Add(
			html.P().Attr("id", "1").Add("hello world!"),
			html.Elem("input").Attr("value", "test").SelfClosing(),
		)

		html := el.String()

		if html != `<div id="123" class="main"><p id="1">hello world!</p><input value="test" /></div>` {
			t.Error(html)
		}
	})

	t.Run("should render with children pretty", func(t *testing.T) {
		el := html.Elem("div")
		el.Attr("id", "123")
		el.Attr("class", "main")
		el.Add(
			html.P().Attr("id", "1").Add("hello world!"),
			html.Elem("input").Attr("value", "test").SelfClosing(),
		)

		html := el.PrettyString("\t")

		if html != "<div id=\"123\" class=\"main\">\n\t<p id=\"1\">hello world!</p>\n\t<input value=\"test\" />\n</div>" {
			t.Error(html)
		}
	})

	t.Run("should render with children nested pretty", func(t *testing.T) {
		el := html.Elem("div")
		el.Attr("id", "123")
		el.Attr("class", "main")
		el.Add(
			html.P().Attr("id", "1").
				Add("hello world!").
				Add(html.Elem("span").Add("hi!")),
			html.Elem("input").Attr("value", "test").SelfClosing(),
		)

		html := el.PrettyString("\t")

		if html != "<div id=\"123\" class=\"main\">\n\t<p id=\"1\">\n\t\thello world!\n\t\t<span>hi!</span>\n\t</p>\n\t<input value=\"test\" />\n</div>" {
			t.Error(html)
		}
	})
}
