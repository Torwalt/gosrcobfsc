package renamer_test

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
	"github.com/stretchr/testify/require"
)

var testFile = `
package main

import (
	"go/types"
	"sync"
)

func asd() {
	numbers := [][]int{}

	wg := sync.WaitGroup{}
	for _, n := range numbers {
		go func() {
			defer wg.Done()

			_ = &types.Config{}
			for _, sp := range n {
				print(sp)
			}
		}()
	}
}
`

func TestFileRenamer(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", testFile, 0)
	require.NoError(t, err)

	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}
	config := &types.Config{
		Importer: importer.ForCompiler(fset, "source", nil),
	}

	path := filepath.Join(moduleName, "cmd")
	_, err = config.Check(path, fset, []*ast.File{f}, info)
	require.NoError(t, err)

	tc := renamer.NewTypeChecker(info)
	ic := renamer.NewImportChecker(f, moduleName)
	fr := renamer.NewFileRenamer(f, ic, tc)

	buf := []byte{}
	bb := bytes.NewBuffer(buf)

	fr.Rename()
	err = format.Node(bb, fset, f)
	require.NoError(t, err)
	out := bb.String()
	fmt.Println(out)
	require.NotEmpty(t, out)
}
