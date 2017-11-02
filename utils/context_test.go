package utils

import (
	"testing"
)

func TestNewContext(t *testing.T) {
	expectedKeys := []string{
		"cwd", "home",
	}

	context := *NewContext(nil)

	for _, name := range expectedKeys {
		if context[name] == nil {
			t.Fatalf("Should have `%s` in context", name)
		}
	}

}
