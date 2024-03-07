package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestConfusingResults(t *testing.T) {
	testRule(t, "confusing-results", &rule.ConfusingResultsRule{})
}
