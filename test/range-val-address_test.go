package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/lint"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestRangeValAddress(t *testing.T) {
	testRule(t, "range-val-address", &rule.RangeValAddress{}, &lint.RuleConfig{})
}
