package obfuscating_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/Torwalt/gosrcobfsc/obfuscating"
	"github.com/stretchr/testify/require"
)

var file = `
package main

import (
    "fmt"
	"github.com/Torwalt/gosrcobfsc/obfuscating"
)

const x = "a"

const (
    y = "b"
    z = "c"
)

var h = "h"

func aFunction(in string) int {
	fmt.Println(in)
	return 0
}

func secondFunc(in string) int {
    x := 1
	return 0
}

type JustAStruct struct {
    APublicField int
    privateField string
}
`

func TestVisit(t *testing.T) {
	content := file

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, 0)
	require.NoError(t, err)

	v := obfuscating.NewVisitor(fset)
	ast.Walk(v, f)

	ns := v.NamedSymbols()

	require.NotEmpty(t, ns)
}
