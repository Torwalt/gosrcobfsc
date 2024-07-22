package obfuscating

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

type Repository []Package

type Package struct {
	fset     *token.FileSet
	name     string
	fullPath string
	pkgMap   map[string]*ast.Package
}

func NewRepository(dirs Dirs) (Repository, error) {
	repo := Repository{}
	for _, dir := range dirs {
		fset := token.NewFileSet()
		pkgMap, err := parser.ParseDir(fset, dir, nil, 0)
		if err != nil {
			return repo, err
		}
		repo = append(repo, Package{
			fset:     fset,
			name:     "",
			fullPath: dir,
			pkgMap:   pkgMap,
		})
	}

	return repo, nil
}

func WriteObfuscated(in []Obfuscated, args Args) error {
	for _, obf := range in {
		// We are in a subdir.
		if obf.Pkg.fullPath != args.Source {
			snk := SinkiFy(args, obf.Pkg.fullPath)
			if err := os.MkdirAll(snk, os.ModePerm); err != nil {
				return err
			}
		}

		for _, pkg := range obf.Pkg.pkgMap {
			for name, file := range pkg.Files {
				snk := SinkiFy(args, name)
				f, err := os.Create(snk)
				if err != nil {
					return err
				}

				if err := format.Node(f, obf.Pkg.fset, file); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
