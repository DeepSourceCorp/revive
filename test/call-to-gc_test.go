package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

// TestCallToGC test call-to-gc rule
func TestCallToGC(t *testing.T) {
	testRule(t, "call-to-gc", &rule.CallToGCRule{})
}
