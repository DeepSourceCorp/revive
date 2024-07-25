package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestRangeValInClosure(t *testing.T) {
	testRule(t, "range-val-in-closure", &rule.RangeValInClosureRule{}, &lint.RuleConfig{})
}
