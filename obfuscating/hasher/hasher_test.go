package hasher_test

import (
	"fmt"
	"testing"

	"github.com/Torwalt/gosrcobfsc/obfuscating/hasher"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "JustAStruct",
			expected: "bhbebfhdagfecigcbfcddibjbaehjbgaeeeachiicacadcbidfefhdaffdbgbcce",
		},
		{
			input:    "JustAStrucT",
			expected: "cccgfdfggdffedbbfdbeeeiaadfefcabajbdcfjgfccefibeeffceehcddibjaeh",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test with input: %v", tt.input), func(t *testing.T) {
			act := hasher.Hash(tt.input)
			require.Equal(t, tt.expected, act)
		})
	}
}
