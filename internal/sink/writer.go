package sink

import (
	"errors"
	"fmt"
	"go/format"
	"os"
	"path/filepath"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/paths"
	"github.com/Torwalt/gosrcobfsc/internal/repo/gomod"
)

const gomodFilename = "go.mod"

func WriteObfuscated(in obfuscate.ObfuscatedRepository, a args.Args) error {
	for _, obfPkg := range in.Packages {
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

	writeGomod(a.Sink, in.ObfuscatedGomod)

	return nil
}

func writeGomod(sink string, gmd gomod.GoMod) error {
	gmodFilePath := filepath.Join(sink, gomodFilename)
	f, err := os.Create(gmodFilePath)
	if err != nil {
		return err
	}
	_, err = f.WriteString(gmd.String())
	if err != nil {
		return err
	}

	return nil
}
