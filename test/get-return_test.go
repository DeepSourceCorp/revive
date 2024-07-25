package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestGetReturn(t *testing.T) {
	testRule(t, "get-return", &rule.GetReturnRule{})
}
