package hasher_test

import (
	"fmt"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/hasher"
	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "JustAStruct",
			expected: "JhbebfhdagfecigcbfcddibjbaehjbgaeeeachiicacadcbidfefhdaffdbgbcceXXX",
		},
		{
			input:    "JustAStrucT",
			expected: "JccgfdfggdffedbbfdbeeeiaadfefcabajbdcfjgfccefibeeffceehcddibjaehXXX",
		},
		{
			input:    "justAStrucT",
			expected: "ecbadbbdfeaceceaifieccfecchfcgdbbbbfaegfcgiaeifbdbbjebaeaeadddbiXXX",
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Test with input: %v", tt.input), func(t *testing.T) {
			act := hasher.Hash(tt.input)
			require.Equal(t, tt.expected, act)
		})
	}
}
