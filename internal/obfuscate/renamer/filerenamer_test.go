package renamer_test

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
	"github.com/stretchr/testify/require"
)

var testFile = `
package main

import (
    "fmt"
	"github.com/Torwalt/gosrcobfsc/obfuscating"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
	"go/format"
)

type ObfuscatedPackage struct {
	Package         *repo.Package
	ObfuscatedPath  renamer.ObfuscatedPath
}

func aFunc(in string) int {
    if err := format.Node(nil, nil, nil); err != nil {
        return err
    }

    x := 1

    ok := true
    if !ok {
    }

    y := []string{}

	return 0
}
`

func TestFileRenamer(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", testFile, 0)
	require.NoError(t, err)

	ic := renamer.NewImportChecker(f, moduleName)
	fr := renamer.NewFileRenamer(f, ic)

	buf := []byte{}
	bb := bytes.NewBuffer(buf)

	fr.Rename()
	err = format.Node(bb, fset, f)
	require.NoError(t, err)
	out := bb.String()
	fmt.Println(out)
	require.NotEmpty(t, out)
}
