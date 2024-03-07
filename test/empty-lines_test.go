package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

// TestEmptyLines rule.
func TestEmptyLines(t *testing.T) {
	testRule(t, "empty-lines", &rule.EmptyLinesRule{})
}
