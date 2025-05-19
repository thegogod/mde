package tokens_test

import (
	"testing"

	"github.com/thegogod/mde/tokens"
)

var src = `# Test

1. a
2. b
3. c

@if (true) {
	4. d
	5. e
}`

func TestScanner(t *testing.T) {
	t.Run("should scan", func(t *testing.T) {
		scanner := tokens.NewScanner(
			"README.md",
			[]byte(src),
		)

		for {
			token, err := scanner.Scan()

			if token == nil || token.Kind == tokens.Eof {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%s => %s", token.Kind, token.String())
		}
	})
}
