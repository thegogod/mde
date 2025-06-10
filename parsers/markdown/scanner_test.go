package markdown_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/parsers/markdown"
)

func TestScanner(t *testing.T) {
	t.Run("should scan", func(t *testing.T) {
		t.SkipNow()
		data, err := os.ReadFile(filepath.Join("..", "testcases", "links.md"))

		if err != nil {
			t.Fatal(err)
		}

		scanner := markdown.NewScanner(data)

		for {
			token, err := scanner.Scan()

			if token == nil || token.Kind() == core.Eof {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%s => %s", token.KindString(), token.String())
		}
	})
}
