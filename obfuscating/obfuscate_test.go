package obfuscating_test

import (
	"github.com/Torwalt/gosrcobfsc/obfuscating"
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
	obfuscated, err := obfuscating.Obfuscate(string(file))
	require.NoError(t, err)
	require.NotEmpty(t, obfuscated)
}
