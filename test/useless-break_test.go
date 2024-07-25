package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

// UselessBreak rule.
func TestUselessBreak(t *testing.T) {
	testRule(t, "useless-break", &rule.UselessBreak{})
}
