package markdown_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/thegogod/mde/parsers/markdown"
)

func TestParser(t *testing.T) {
	RunDir(t, filepath.Join("..", "..", "testcases"), func(t *testing.T, data []byte) {
		parser := markdown.New()
		node, err := parser.Parse(data)

		if err != nil {
			t.Fatal(err)
		}

		t.Log(node.PrettyString("  "))
	})
}

func RunDir(t *testing.T, path string, fn func(t *testing.T, data []byte)) error {
	entries, err := os.ReadDir(path)

	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())

		if entry.IsDir() {
			RunDir(t, entryPath, fn)
		} else if strings.HasSuffix(entry.Name(), ".md") {
			RunFile(t, entryPath, fn)
		}
	}

	return nil
}

func RunFile(t *testing.T, path string, fn func(t *testing.T, data []byte)) error {
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	t.Run(filepath.Base(path), func(t *testing.T) {
		fn(t, data)
	})

	return nil
}
