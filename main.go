package certen

import (
	"log"
	"os"
	"certen/certen"
)

var run = certen.Run

func main() {
	log.SetFlags(0)
	log.SetPrefix("certen: ")
	log.SetOutput(os.Stderr)
	run()
}
