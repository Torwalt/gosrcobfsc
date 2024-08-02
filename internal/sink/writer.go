package sink

import (
	"errors"
	"fmt"
	"go/format"
	"os"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/paths"
)

func WriteObfuscated(in obfuscate.ObfuscatedRepository, a args.Args) error {
	for _, obfPkg := range in {
		// Make dir of package
		snk := args.SinkiFy(a, obfPkg.ObfuscatedPath.Full())
		if err := os.MkdirAll(snk, os.ModePerm); err != nil {
			return err
		}

		for _, astPkg := range obfPkg.Package.PkgMap {
			for name, file := range astPkg.Files {
				nrp := paths.NonRootPath(name, a.Source)
				obfFilePath, ok := obfPkg.ObfuscatedFiles[nrp]
				if !ok {
					return errors.New(fmt.Sprintf("Missing obfuscated filepath for filepath: %v", name))
				}

				snk := args.SinkiFy(a, obfFilePath.Full())
				f, err := os.Create(snk)
				if err != nil {
					return err
				}

				if err := format.Node(f, obfPkg.Package.Fset, file); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
