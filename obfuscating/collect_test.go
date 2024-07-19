package obfuscating_test

import (
	"testing"

	"github.com/Torwalt/gosrcobfsc/obfuscating"
	"github.com/stretchr/testify/require"
)

var thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"

func TestCollectDirs(t *testing.T) {
	d, err := obfuscating.CollectDirs(thisRepoFullPath)
	require.NoError(t, err)
	require.NotEmpty(t, d)
	require.NotContains(t, d, "/home/ada/repos/gosrcobfsc/.git")
}
