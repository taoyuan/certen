package caddy

import (
	"certen"
	"certen/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func NewCaddy() (*Caddy, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return New(filepath.Join(wd, "../../fixtures/caddy"))
}

func TestNew(t *testing.T) {
	caddy, err := NewCaddy()
	if err != nil {
		t.Fatal(err)
	}
	if caddy.dir == "" {
		t.Fatal("\"dir\" of caddy instance is empty")
	}
}

func TestCaddy_FindDomains(t *testing.T) {
	caddy, err := NewCaddy()
	if err != nil {
		t.Fatal(err)
	}

	var domains []*certen.Domain

	domains, err = caddy.FindDomains("*")
	if err != nil {
		t.Fatal(err)
	}
	if len(domains) != 2 {
		t.Fatalf("There are %d domains found, but expected %d domains", len(domains), 2)
	}

	domains, err = caddy.FindDomains("abc.com,example.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(domains) != 2 {
		t.Fatalf("There are %d domains found, but expected %d domains", len(domains), 2)
	}

	domains, err = caddy.FindDomains("*.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(domains) != 2 {
		t.Fatalf("There are %d domains found, but expected %d domains", len(domains), 2)
	}

	domains, err = caddy.FindDomains("abc.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(domains) != 1 {
		t.Fatalf("There are %d domains found, but expected %d domains", len(domains), 1)
	}
}

func TestCaddy_Export(t *testing.T) {
	caddy, err := NewCaddy()
	if err != nil {
		t.Fatal(err)
	}

	// Test combining
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = caddy.ExportByName("*", dir, true)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range []string{
		filepath.Join(dir, "abc.com.pem"),
		filepath.Join(dir, "example.com.pem"),
	} {
		if !utils.Exists(file) {
			t.Fatal("Should have " + file + " file created, but it not exists")
		}
	}

	// Test without combining
	dir, err = ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = caddy.ExportByName("*", dir, false)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range []string{
		filepath.Join(dir, "abc.com.crt"),
		filepath.Join(dir, "abc.com.key"),
		filepath.Join(dir, "example.com.crt"),
		filepath.Join(dir, "example.com.key"),
	} {
		if !utils.Exists(file) {
			t.Fatal("Should have " + file + " file created, but it not exists")
		}
	}

}
