package main

import (
	"gosrcobfsc/obfusc"
	"log"
)

func main() {
	// TODO
	// - Accept args
	// - Read file
	fileContent := ""

	if err := run(fileContent); err != nil {
		log.Fatalf("%v", err)
	}
}

func run(content string) error {
	out, err := obfusc.Obfuscate(content)
	if err != nil {
		return err
	}
	log.Printf("Obfuscated file: %v", out)

	return nil
}
