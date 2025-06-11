package markdown_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/thegogod/mde"
)

func TestMarkdown(t *testing.T) {
	RunDir(t, filepath.Join("testcases"), func(t *testing.T, md []byte, html []byte) {
		parser := mde.New()
		node, err := parser.Parse(md)

		if err != nil {
			t.Fatal(err)
		}

		pretty := node.PrettyBytes("  ")

		if html == nil {
			t.Log(string(pretty))
			return
		}

		if !bytes.Equal(pretty, html) {
			t.Logf("expected: %v", string(html))
			t.Logf("received: %v", string(pretty))
			t.FailNow()
		}
	})
}

func RunDir(t *testing.T, path string, fn func(t *testing.T, md []byte, html []byte)) error {
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

func RunFile(t *testing.T, path string, fn func(t *testing.T, md []byte, html []byte)) error {
	html, _ := os.ReadFile(strings.Replace(path, ".md", ".html", 1))
	md, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	t.Run(filepath.Base(path), func(t *testing.T) {
		fn(t, md, html)
	})

	return nil
}
