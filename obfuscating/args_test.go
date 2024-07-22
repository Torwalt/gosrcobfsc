package obfuscating_test

import (
	"path/filepath"
	"testing"

	"github.com/Torwalt/gosrcobfsc/obfuscating"
	"github.com/stretchr/testify/require"
)

func TestSinkiFy(t *testing.T) {
	args, err := obfuscating.NewArgs(&moduleName, &thisRepoFullPath, &sink)
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

		sinkyfied := obfuscating.SinkiFy(args, sourceFilePath)
		require.Equal(t, tt.expected, sinkyfied)
	}
}
