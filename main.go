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
	flag.Parse()

	args, err := obfuscating.NewArgs(moduleNameFlag, sourceFlag, sinkFlag)
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("%v", err)
	}

	if err := run(args); err != nil {
		log.Fatalf("%v", err)
	}
}

func run(args obfuscating.Args) error {
	_, err := obfuscating.Obfuscate(args)
	if err != nil {
		return err
	}

	log.Printf("Successfully obfuscated %v and wrote result into %v", args.Source, args.Sink)

	return nil
}
