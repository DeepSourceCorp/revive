package test

import (
	"testing"

	"github.com/DeepSourceCorp/revive/rule"
)

func TestUnexportednaming(t *testing.T) {
	testRule(t, "unexported-naming", &rule.UnexportedNamingRule{})
}
