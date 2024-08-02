package main

import (
	"flag"
	"log"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/Torwalt/gosrcobfsc/internal/repo/gitignore"
	"github.com/Torwalt/gosrcobfsc/internal/sink"
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
	gi, err := gitignore.NewFromFilePath(a.Source)
	if err != nil {
		return err
	}

	dirs, err := repo.CollectDirs(a.Source, repo.FilterFuncWithGitIgnore(gi, a.Source))
	if err != nil {
		return err
	}

	rpo, err := repo.NewRepository(dirs, a.Source, a.ModuleName)
	if err != nil {
		return err
	}

	or, err := obfuscate.Obfuscate(rpo)
	if err != nil {
		return err
	}

	err = sink.WriteObfuscated(or, a)
	if err != nil {
		return err
	}

	log.Printf("Successfully obfuscated %v and wrote result into %v", a.Source, a.Sink)

	return nil
}
