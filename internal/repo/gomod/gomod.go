package gomod

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const gomodFilename = "go.mod"

type GoMod struct {
	Version    string
	ModuleName string
}

func NewGoMod(in string) (GoMod, error) {
	if in == "" {
		return GoMod{}, errors.New("no content in go.mod")
	}

	lineSplit := strings.Split(in, "\n")
	filtered := filterWhitespace(lineSplit)

	if len(filtered) < 2 {
		return GoMod{}, errors.New(fmt.Sprintf("given go.mod content is incomplete: %v", filtered))
	}

	moduleLine := filtered[0]
	// Expect: module github.com/Torwalt/gosrcobfsc
	moduleSplitted := strings.Split(moduleLine, " ")

	if len(moduleSplitted) < 2 {
		return GoMod{}, errors.New(fmt.Sprintf("missing or incomplete module line: %v", moduleSplitted))
	}

	gomod := GoMod{
		ModuleName: moduleSplitted[1],
	}

	versionLine := filtered[1]
	// Expect: go 1.22.4
	versionSplitted := strings.Split(versionLine, " ")

	if len(versionSplitted) < 2 {
		return GoMod{}, errors.New(fmt.Sprintf("missing or incomplete version line: %v", versionSplitted))
	}

	gomod.Version = versionSplitted[1]

	return gomod, nil
}

func NewGoModFromPath(path string) (GoMod, error) {
	gomodPath := filepath.Join(path, gomodFilename)
	content, err := os.ReadFile(gomodPath)
	if err != nil {
		return GoMod{}, errors.New("no gomod file found in path")
	}

	return NewGoMod(string(content))
}

func filterWhitespace(in []string) []string {
	out := []string{}
	for _, i := range in {
		trimmed := strings.TrimSpace(i)
		if trimmed == "" {
			continue
		}
		out = append(out, trimmed)
	}

	return out
}
