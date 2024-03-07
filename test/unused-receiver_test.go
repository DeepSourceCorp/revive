package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestUnusedReceiver(t *testing.T) {
	testRule(t, "unused-receiver", &rule.UnusedReceiverRule{})
}
