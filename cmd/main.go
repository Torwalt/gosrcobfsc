package main

import (
	"flag"
	"log"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
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
	dirs, err := repo.CollectDirs(a.Source)
	if err != nil {
		return err
	}

	rpo, err := repo.NewRepository(dirs)
	if err != nil {
		return err
	}

	rpo, err = obfuscate.Obfuscate(rpo)
	if err != nil {
		return err
	}

	err = repo.WriteObfuscated(rpo, a)
	if err != nil {
		return err
	}

	log.Printf("Successfully obfuscated %v and wrote result into %v", a.Source, a.Sink)

	return nil
}
