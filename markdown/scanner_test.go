package markdown_test

import (
	"testing"

	"github.com/thegogod/mde/markdown"
)

var src = `# Title

my description...

## Sub Title

**_some more text..._**

1. __a__
2. b
3. c

- test
- ing

> a test quote

---

[google](https://google.com)

![IMAGE!!!](https://google.com/images/batman.jpg)

`

func TestScanner(t *testing.T) {
	t.Run("should scan", func(t *testing.T) {
		scanner := markdown.NewScanner([]byte(src))

		for {
			token, err := scanner.Scan()

			if token == nil || markdown.TokenKind(token.GetTokenKind()) == markdown.Eof {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%s => %s", markdown.TokenKind(token.GetTokenKind()), token.String())
		}
	})
}
