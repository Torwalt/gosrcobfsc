package args_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/stretchr/testify/require"
)

var (
	thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"
	moduleName       = "github.com/Torwalt/gosrcobfsc"
	sink             = filepath.Join(thisRepoFullPath, "tests")
)

func TestSinkiFy(t *testing.T) {
	a, err := args.NewArgs(&moduleName, &thisRepoFullPath, &sink)
	require.NoError(t, err)

	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "main.go",
			expected: "/home/ada/repos/gosrcobfsc/tests/main.go",
		},
		{
			input:    "/obfuscating/args.go",
			expected: "/home/ada/repos/gosrcobfsc/tests/obfuscating/args.go",
		},
	}

	for _, tt := range tests {
		sourceFilePath := filepath.Join(thisRepoFullPath, tt.input)

		sinkyfied := args.SinkiFy(a, sourceFilePath)
		require.Equal(t, tt.expected, sinkyfied)
	}
}
