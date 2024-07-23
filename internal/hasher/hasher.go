package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type Hashed = string

// NOTE: Maybe I will need to control the hash length because that might cause
// problems with either go syntax or later on directory/package names.
func Hash(in string) Hashed {
	hash := sha256.Sum256([]byte(in))

	hexHash := hex.EncodeToString(hash[:])

	letters := hexToLetters(hexHash)

	return Hashed(letters)
}

func hexToLetters(hexString string) string {
	var out strings.Builder

	for _, r := range hexString {
		var ru rune
		switch r {
		case '0':
			ru = 'a'
		case '1':
			ru = 'b'
		case '2':
			ru = 'c'
		case '3':
			ru = 'd'
		case '4':
			ru = 'e'
		case '5':
			ru = 'f'
		case '6':
			ru = 'g'
		case '7':
			ru = 'h'
		case '8':
			ru = 'i'
		case '9':
			ru = 'j'
		default:
			ru = r
		}

		out.WriteRune(ru)
	}

	return out.String()
}