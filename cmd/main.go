package main

import (
	"flag"
	"log"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
)

func main() {
	moduleNameFlag := flag.String("moduleName", "", "The name of the module (top of go.mod).")
	sourceFlag := flag.String("source", "", "The full path of the source repository.")
	sinkFlag := flag.String("sink", "", "The full path where to write obfuscated directory.")
	flag.Parse()

	args, err := args.NewArgs(moduleNameFlag, sourceFlag, sinkFlag)
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("%v", err)
	}

	if err := run(args); err != nil {
		log.Fatalf("%v", err)
	}
}

func run(a args.Args) error {
	_, err := obfuscate.Obfuscate(a)
	if err != nil {
		return err
	}

	log.Printf("Successfully obfuscated %v and wrote result into %v", a.Source, a.Sink)

	return nil
}
