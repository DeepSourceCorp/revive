package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestUseAny(t *testing.T) {
	testRule(t, "use-any", &rule.UseAnyRule{})
}
