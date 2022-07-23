package main

import (
	"log"
	"os"

	"github.com/koron-go/z80op/opcode"
)

// Dump decode layer to STDOUT.

func main() {
	err := opcode.DumpDecodeLayer(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
