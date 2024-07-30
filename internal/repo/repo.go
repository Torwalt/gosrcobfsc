package repo

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"

	"github.com/Torwalt/gosrcobfsc/internal/args"
)

type Repository []Package

type Package struct {
	fset     *token.FileSet
	fullPath string
	PkgMap   map[string]*ast.Package
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
			fullPath: dir,
			PkgMap:   pkgMap,
		})
	}

	return repo, nil
}

func WriteObfuscated(in []Package, a args.Args) error {
	for _, pkg := range in {
		// We are in a subdir.
		if pkg.fullPath != a.Source {
			snk := args.SinkiFy(a, pkg.fullPath)
			if err := os.MkdirAll(snk, os.ModePerm); err != nil {
				return err
			}
		}

		for _, astPkg := range pkg.PkgMap {
			for name, file := range astPkg.Files {
				snk := args.SinkiFy(a, name)
				f, err := os.Create(snk)
				if err != nil {
					return err
				}

				if err := format.Node(f, pkg.fset, file); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
