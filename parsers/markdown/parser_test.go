package markdown_test

import (
	"testing"

	"github.com/thegogod/mde/parsers/markdown"
)

var parserSrc = `# *Title*

my description...

## Sub Title

**_some more text..._**  ` +
	"```some *inline* code...```"

func TestParser(t *testing.T) {
	t.Run("should parse", func(t *testing.T) {
		parser := markdown.Parser{}
		node, err := parser.Parse([]byte(parserSrc))

		if err != nil {
			t.Fatal(err)
		}

		value, err := node.Eval()

		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(value.Bytes()))
	})
}
