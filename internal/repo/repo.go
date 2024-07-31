package repo

import (
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/hasher"
	"github.com/Torwalt/gosrcobfsc/internal/paths"
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
			hashed := hashPath(a.Source, pkg.fullPath)
			snk := args.SinkiFy(a, hashed)
			if err := os.MkdirAll(snk, os.ModePerm); err != nil {
				return err
			}
		}

		for _, astPkg := range pkg.PkgMap {
			for name, file := range astPkg.Files {
				hashed := hashPath(a.Source, name)
				snk := args.SinkiFy(a, hashed)
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

func hashPath(root, path string) string {
	nonRootPath, _ := strings.CutPrefix(path, root)
	ext := filepath.Ext(nonRootPath)
	split := paths.SplitAndFilter(nonRootPath, paths.FilterEmpty)

	hashed := make([]string, 0, len(split))
	for idx, p := range split {
		hash := hasher.Hash(p)
		if idx == len(split)-1 {
			// Append the .go extension to hashed path.
			hash += ext
		}
		hashed = append(hashed, hash)
	}

	return filepath.Join(hashed...)
}
