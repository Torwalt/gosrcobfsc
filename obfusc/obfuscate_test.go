package obfusc_test

import (
	"gosrcobfsc/obfusc"
	"testing"

	"github.com/stretchr/testify/require"
)

var file = `
package main

import "fmt"

func aFunction(in string) int {
	fmt.Println(in)
	return 0
}
`

func TestObfuscate(t *testing.T) {
	// fileContent, err := os.ReadFile("./testfile.go.test")
	// require.NoError(t, err)

	obfuscated, err := obfusc.Obfuscate(string(file), 1337)
	require.NoError(t, err)
	require.NotEmpty(t, obfuscated)
}
