package tokens_test

import (
	"testing"

	"github.com/thegogod/mde/tokens"
)

var src = `# Hello World
this is a test`

func TestScanner(t *testing.T) {
	t.Run("should scan", func(t *testing.T) {
		scanner := tokens.NewScanner(
			"README.md",
			[]byte(src),
		)

		token, err := scanner.Scan()

		if err != nil {
			t.Fatal(err)
		}

		t.Log(token)
	})
}
