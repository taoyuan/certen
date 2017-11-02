package certen

import (
	"errors"
)

type Domain struct {
	Name string
	Dir  string
}

type Provider interface {
	FindDomains(pattern string) ([]*Domain, error)
	Export(domains []*Domain, dest string, assemble bool) ([]string, error)
	ExportByName(name string, dest string, assemble bool) ([]string, error)
}

type ProviderCreator func(dir string) (Provider, error)

var creators = make(map[string]ProviderCreator)

func RegisterProvider(name string, creator ProviderCreator) error {
	if name == "" {
		panic("plugin must have a name")
	}
	if _, dup := creators[name]; dup {
		panic("provider named " + name + " already registered")
	}

	creators[name] = creator
	return nil
}

func HasProvider(name string) bool {
	return creators[name] != nil
}

func CreateProvider(name string, dir string) (Provider, error) {
	if !HasProvider(name) {
		return nil, errors.New("Unknown provider name " + name)
	}
	return creators[name](dir)
}
