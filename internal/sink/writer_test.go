package sink_test

import (
	"testing"

	"github.com/Torwalt/gosrcobfsc/internal/args"
	"github.com/Torwalt/gosrcobfsc/internal/obfuscate"
	"github.com/Torwalt/gosrcobfsc/internal/repo"
	"github.com/Torwalt/gosrcobfsc/internal/repo/gitignore"
	"github.com/Torwalt/gosrcobfsc/internal/sink"
	"github.com/stretchr/testify/require"
)

var (
	thisRepoFullPath = "/home/ada/repos/gosrcobfsc/"
	moduleName       = "github.com/Torwalt/gosrcobfsc"
)

func TestWriteObfuscated(t *testing.T) {
	snk := t.TempDir()
	args, err := args.NewArgs(&moduleName, &thisRepoFullPath, &snk)
	require.NoError(t, err)

	gi, err := gitignore.NewFromFilePath(thisRepoFullPath)
	require.NoError(t, err)

	dirs, err := repo.CollectDirs(args.Source, repo.FilterFuncWithGitIgnore(gi, args.Source))
	require.NoError(t, err)

	rpo, err := repo.NewRepository(dirs, args.Source)
	require.NoError(t, err)
	require.NotEmpty(t, rpo)

	or, err := obfuscate.Obfuscate(rpo)
	require.NoError(t, err)
	require.NotEmpty(t, or)

	err = sink.WriteObfuscated(or, args)
	require.NoError(t, err)
}
