package obfuscate

import (
	"go/ast"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
)

/*

Hard part: module package must be renamed, too, but external packages cannot be renamed.
Super hard part: What if code has custom private packages imported??

*/

/*
TODO:
- pkg and directory renaming
- lookup of packages and their paths to change imports
- args controlling
- writing resulting repo
- walking a whole repo not just file
*/

type Obfuscated struct {
	Pkg repo.Package
	Ns  NamedSymbols
}

func Obfuscate(args args.Args) ([]Obfuscated, error) {
	dirs, err := repo.CollectDirs(args.Source)
	if err != nil {
		return []Obfuscated{}, err
	}

	rpo, err := repo.NewRepository(dirs)
	if err != nil {
		return []Obfuscated{}, err
	}

	// I.e., we need to walk the rpo and mutate the NamedSymbols.
	obf := []Obfuscated{}
	for _, pkgs := range rpo {
		for _, pkg := range pkgs.PkgMap {
			v := NewVisitor()
			ast.Walk(v, pkg)
			obf = append(obf, Obfuscated{
				Pkg: pkgs,
				Ns:  v.NamedSymbols(),
			})
		}
	}

	if err = repo.WriteObfuscated(rpo, args); err != nil {
		return obf, err
	}

	return obf, nil
}
