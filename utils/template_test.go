package utils

import (
	"log"
	"testing"
)

func TestExecute(t *testing.T) {
	context := Context{"name": "Tom"}
	answer, err := Execute("{{.name}}", &context)
	if err != nil {
		t.Error("template execute failure", err)
	}
	if answer != context["name"] {
		log.Fatalf("template execute with wrong answer: %s, expected: %s", answer, context["name"])
	}
}
