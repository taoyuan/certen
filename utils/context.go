package utils

import (
	"log"
	"os"
	"os/user"
)

type Context map[string]interface{}

var DefaultContext = Context{}

func init() {
	// init cwd
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	DefaultContext["wd"] = wd
	DefaultContext["cwd"] = wd

	// init home
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	DefaultContext["home"] = usr.HomeDir
}

func NewContext(data *Context) *Context {
	answer := Context{}
	for k, v := range DefaultContext {
		answer[k] = v
	}
	if data != nil {
		c := *data
		for k, v := range c {
			answer[k] = v
		}
	}
	return &answer
}
