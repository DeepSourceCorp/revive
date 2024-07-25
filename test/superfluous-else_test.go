package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/internal/ifelse"
	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

// TestSuperfluousElse rule.
func TestSuperfluousElse(t *testing.T) {
	testRule(t, "superfluous-else", &rule.SuperfluousElseRule{})
	testRule(t, "superfluous-else-scope", &rule.SuperfluousElseRule{}, &lint.RuleConfig{Arguments: []any{ifelse.PreserveScope}})
}
