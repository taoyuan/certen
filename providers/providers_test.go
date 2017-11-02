package providers

import (
	"certen"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateProvider(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	provider, err := certen.CreateProvider("caddy", filepath.Join(wd, "../fixtures/caddy"))
	if err != nil {
		t.Fatal(err)
	}
	if provider == nil {
		t.Fatal("Should create provider caddy")
	}
}
