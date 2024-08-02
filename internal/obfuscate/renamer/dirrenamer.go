package renamer

import (
	"path/filepath"
	"strings"

	"github.com/Torwalt/gosrcobfsc/internal/hasher"
	"github.com/Torwalt/gosrcobfsc/internal/paths"
)

const (
	goExt      = ".go"
	testSuffix = "_test"
)

// RenamePackage returns an ObfuscatedPath. It is expected that the path is a
// relative part that is relevant to the repository.
// Good: internal/obfuscate/renamer/dirrenamer.go
// Bad: /home/user/gosrcobfsc/internal/obfuscate/renamer/dirrenamer.go
func RenamePackage(path string) ObfuscatedPath {
	if path == "" {
		return ObfuscatedPath{}
	}

	p := newPath(path)
	op := newObfuscatedPath(p)
	return op
}

type ObfuscatedPath struct {
	Path     string
	Filename string
}

func (op ObfuscatedPath) Full() string {
	return filepath.Join(op.Path, op.Filename)
}

func newObfuscatedPath(p path) ObfuscatedPath {
	hashedParts := []string{}
	for _, p := range p.parts {
		hashedParts = append(hashedParts, hasher.Hash(p))
	}
	hps := filepath.Join(hashedParts...)

	op := ObfuscatedPath{
		Path:     hps,
		Filename: "",
	}

	if p.filename == "" {
		return op
	}

	hpath := hasher.Hash(p.filename)
	suffix := goExt
	if p.isTest {
		suffix = testSuffix + suffix
	}

	op.Filename = hpath + suffix

	return op
}

type path struct {
	parts    []string
	filename string
	isTest   bool
}

func newPath(in string) path {
	dir, filename := filepath.Split(in)
	parts := paths.SplitAndFilter(dir, paths.FilterEmpty)

	p := path{
		parts:    parts,
		filename: "",
		isTest:   false,
	}

	if filename == "" {
		return p
	}

	if filepath.Ext(filename) != goExt {
		// We assume its just the last dir element.
		p.parts = append(p.parts, filename)

		return p
	}

	p.filename = filename
	p.isTest = strings.Contains(filename, testSuffix)

	return p
}
