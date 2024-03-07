package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

// ConstantLogicalExpr rule.
func TestConstantLogicalExpr(t *testing.T) {
	testRule(t, "constant-logical-expr", &rule.ConstantLogicalExprRule{})
}
