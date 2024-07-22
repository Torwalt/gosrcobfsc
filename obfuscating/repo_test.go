package obfuscating_test

import (
	"testing"

	"github.com/Torwalt/gosrcobfsc/obfuscating"
	"github.com/stretchr/testify/require"
)

func TestWriteObfuscated(t *testing.T) {
	args, err := obfuscating.NewArgs(&moduleName, &thisRepoFullPath, &sink)
	require.NoError(t, err)

	obf, err := obfuscating.Obfuscate(args)
	require.NoError(t, err)

	err = obfuscating.WriteObfuscated(obf, args)
	require.NoError(t, err)
}
