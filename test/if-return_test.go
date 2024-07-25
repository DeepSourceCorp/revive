package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

// TestIfReturn rule.
func TestIfReturn(t *testing.T) {
	testRule(t, "if-return", &rule.IfReturnRule{})
}
