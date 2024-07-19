package main

import (
	"flag"
	"log"

	"github.com/Torwalt/gosrcobfsc/obfuscating"
)

func main() {
	moduleNameFlag := flag.String("moduleName", "", "The name of the module (top of go.mod).")
	sourceFlag := flag.String("source", "", "The full path of the source repository.")
	sinkFlag := flag.String("sink", "", "The full path where to write obfuscated directory.")
	args, err := obfuscating.NewArgs(moduleNameFlag, sourceFlag, sinkFlag)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := run(args); err != nil {
		log.Fatalf("%v", err)
	}
}

func run(args obfuscating.Args) error {
	out, err := obfuscating.Obfuscate(args)
	if err != nil {
		return err
	}
	log.Printf("Obfuscated file: %v", out)

	return nil
}
