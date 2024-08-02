package renamer_test

import (
	"fmt"
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/obfuscate/renamer"
	"github.com/stretchr/testify/assert"
)

func TestRenamePackage(t *testing.T) {
	tsts := []struct {
		path              string
		expObfuscatedPath renamer.ObfuscatedPath
	}{
		{
			path: "internal/obfuscate/renamer/dirrenamer.go",
			expObfuscatedPath: renamer.ObfuscatedPath{
				Path:     "dbedccbdadacfhbgaiefeaiecaccgicdffcaecgjhgddfecfeffcicjgfgbcafefXXX/addcfeacjgjcdjfdhdedfbfjdbighaihhfjfagafecbbajcdhffdcbaigaicaefjXXX/accibdfeeiajhdchiccgfhadeceecedjjfaegcgbfjaebacefijaeggabieeffaaXXX",
				Filename: "cbcdgggbaeficdgbciajaehcehbebjbefagebhgfdeeiecdadcbcfiieegeagbddXXX.go",
			},
		},
		{
			path: "internal/obfuscate/renamer",
			expObfuscatedPath: renamer.ObfuscatedPath{
				Path:     "dbedccbdadacfhbgaiefeaiecaccgicdffcaecgjhgddfecfeffcicjgfgbcafefXXX/addcfeacjgjcdjfdhdedfbfjdbighaihhfjfagafecbbajcdhffdcbaigaicaefjXXX/accibdfeeiajhdchiccgfhadeceecedjjfaegcgbfjaebacefijaeggabieeffaaXXX",
				Filename: "",
			},
		},
		{
			path: "internal/obfuscate/renamer/dirrenamer_test.go",
			expObfuscatedPath: renamer.ObfuscatedPath{
				Path:     "dbedccbdadacfhbgaiefeaiecaccgicdffcaecgjhgddfecfeffcicjgfgbcafefXXX/addcfeacjgjcdjfdhdedfbfjdbighaihhfjfagafecbbajcdhffdcbaigaicaefjXXX/accibdfeeiajhdchiccgfhadeceecedjjfaegcgbfjaebacefijaeggabieeffaaXXX",
				Filename: "bcccbeccbafaeiagejacdaiafhidbhbiegdabggbbidaehifgfgcaedhdfdaidaiXXX_test.go",
			},
		},
		{
			path:              "",
			expObfuscatedPath: renamer.ObfuscatedPath{},
		},
	}

	for _, tt := range tsts {
		t.Run(fmt.Sprintf("Test with path: %v", tt.path), func(t *testing.T) {
			op := renamer.RenamePackage(tt.path)
			assert.Equal(t, tt.expObfuscatedPath, op)
		})
	}
}
