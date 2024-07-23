package obfuscating

import (
	"go/ast"
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
	Pkg Package
	Ns  NamedSymbols
}

func Obfuscate(args Args) ([]Obfuscated, error) {
	dirs, err := CollectDirs(args.Source)
	if err != nil {
		return []Obfuscated{}, err
	}

	repo, err := NewRepository(dirs)
	if err != nil {
		return []Obfuscated{}, err
	}

	obf := []Obfuscated{}
	for _, pkgs := range repo {
		for _, pkg := range pkgs.pkgMap {
			v := NewVisitor()
			ast.Walk(v, pkg)
			obf = append(obf, Obfuscated{
				Pkg: pkgs,
				Ns:  v.NamedSymbols(),
			})
		}
	}

	if err = WriteObfuscated(obf, args); err != nil {
		return obf, err
	}

	return obf, nil
}
