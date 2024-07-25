package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestWaitGroupByValue(t *testing.T) {
	testRule(t, "waitgroup-by-value", &rule.WaitGroupByValueRule{})
}
