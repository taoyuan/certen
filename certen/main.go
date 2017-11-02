package main

import (
	"certen/certen/certenmain"
	"log"
	"os"
)

var run = certenmain.Run

func main() {
	log.SetFlags(0)
	log.SetPrefix("certen: ")
	log.SetOutput(os.Stderr)
	run()
}
