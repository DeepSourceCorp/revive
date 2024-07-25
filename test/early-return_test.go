package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/internal/ifelse"
	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

// TestEarlyReturn tests early-return rule.
func TestEarlyReturn(t *testing.T) {
	testRule(t, "early-return", &rule.EarlyReturnRule{})
	testRule(t, "early-return-scope", &rule.EarlyReturnRule{}, &lint.RuleConfig{Arguments: []any{ifelse.PreserveScope}})
}
