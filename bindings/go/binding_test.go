package tree_sitter_vala_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-vala"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_vala.Language())
	if language == nil {
		t.Errorf("Error loading Vala grammar")
	}
}
