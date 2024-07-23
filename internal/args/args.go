package args

import (
	"errors"
	"path/filepath"
	"strings"
)

type Args struct {
	// The module name, e.g. for this repo github.com/Torwalt/gosrcobfsc. We
	// need this to identify import paths and pkg names that must be changed.
	// Import paths of external pkgs must not be touched.
	ModuleName string
	// Full file path to the to be obfuscated repo.
	Source string
	// Directory where the new repo should be created.
	Sink string
}

func NewArgs(moduleName, source, sink *string) (Args, error) {
	if moduleName == nil || *moduleName == "" {
		return Args{}, errors.New("moduleName cant be empty")
	}

	if source == nil || *source == "" {
		return Args{}, errors.New("source cant be empty")
	}

	if sink == nil || *sink == "" {
		return Args{}, errors.New("sink cant be empty")
	}

	return Args{
		ModuleName: *moduleName,
		Source:     *source,
		Sink:       *sink,
	}, nil
}

func SinkiFy(in Args, sourceFilePath string) string {
	fp, _ := strings.CutPrefix(sourceFilePath, in.Source)

	return filepath.Join(in.Sink, fp)
}
