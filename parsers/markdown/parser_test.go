package markdown_test

import (
	"testing"

	"github.com/thegogod/mde/parsers/markdown"
)

var parserSrc = `# *Title*

my description...

## Sub Title

**_some more text..._**  ` +
	"```some *inline* code...```" +
	`> a
> > test
> c

[testing123](https://www.google.com)
![an image](https://pasa.org/wp-content/uploads/2021/06/Vervet-Monkey-Foundation-credit-Kyle-.jpg)

**<i>test</i>**

- testing
- a
- list!

1. testing
2. - a
	- nested
	- list
3. ordered
4. list!

`

func TestParser(t *testing.T) {
	t.Run("should parse", func(t *testing.T) {
		parser := markdown.New()
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
