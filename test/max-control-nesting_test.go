package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestMaxControlNesting(t *testing.T) {
	testRule(t, "max-control-nesting", &rule.MaxControlNestingRule{}, &lint.RuleConfig{
		Arguments: []any{int64(2)}},
	)
}
