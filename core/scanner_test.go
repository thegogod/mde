package core_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thegogod/mde/core"
	"github.com/thegogod/mde/extensions/frontmatter"
	"github.com/thegogod/mde/extensions/markdown"
)

func TestScanner(t *testing.T) {
	t.Run("should scan", func(t *testing.T) {
		t.SkipNow()
		data, err := os.ReadFile(filepath.Join("..", "extensions", "markdown", "testcases", "inline_html.md"))

		if err != nil {
			t.Fatal(err)
		}

		scanner := core.NewScanner(data).
			Extend(markdown.New()).
			Extend(frontmatter.New())

		for {
			token, err := scanner.Scan()

			if token == nil || token.Kind() == core.Eof {
				break
			}

			if err != nil {
				t.Fatal(err)
			}
		}
	})
}
