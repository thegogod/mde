package markdown_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thegogod/mde/parsers/markdown"
)

func TestParser(t *testing.T) {
	t.Run("code_block", func(t *testing.T) {
		data, err := os.ReadFile(filepath.Join("..", "..", "testcases", "code_block.md"))

		if err != nil {
			t.Fatal(err)
		}

		parser := markdown.New()
		node, err := parser.Parse(data)

		if err != nil {
			t.Fatal(err)
		}

		value, err := node.Render()

		if err != nil {
			t.Fatal(err)
		}

		t.Log(string(value.Bytes()))
	})
}
