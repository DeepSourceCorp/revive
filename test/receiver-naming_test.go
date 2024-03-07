package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/internal/typeparams"
	"github.com/DeepSourceCorp/revive/rule"
)

func TestReceiverNamingTypeParams(t *testing.T) {
	if !typeparams.Enabled() {
		t.Skip("type parameters are not enabled in the current build environment")
	}
	testRule(t, "receiver-naming-issue-669", &rule.ReceiverNamingRule{})
}
