package tokens_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thegogod/mde/parsers/markdown/tokens"
)

func TestScanner(t *testing.T) {
	t.Run("should scan", func(t *testing.T) {
		t.SkipNow()
		data, err := os.ReadFile(filepath.Join("..", "..", "testcases", "code_block.md"))

		if err != nil {
			t.Fatal(err)
		}

		scanner := tokens.NewScanner(data)

		for {
			token, err := scanner.Scan()

			if token == nil || token.Kind() == tokens.Eof {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			t.Logf("%d => %s", token.Kind(), token.String())
		}
	})
}
