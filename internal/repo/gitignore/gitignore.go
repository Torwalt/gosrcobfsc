package gitignore

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	linebreakSep = "\n"
	gitignore    = ".gitignore"
	gitDir       = ".git"
)

type GitIgnore struct {
	ruleMap             map[string]literalOrExpr
	pathToGitignoreFile string
}

func NewGitIgnore(in, path string) GitIgnore {
	splitted := strings.Split(in, linebreakSep)
	filtered := []string{}
	for _, s := range splitted {
		if s == "" {
			continue
		}
		filtered = append(filtered, s)
	}

	ruleMap := make(ruleMap, len(filtered))
	for _, r := range filtered {
		rg, err := regexp.Compile(r)
		if err != nil {
			ruleMap[r] = literalOrExpr{
				lit:  r,
				expr: nil,
			}
			continue
		}

		ruleMap[r] = literalOrExpr{
			lit:  r,
			expr: rg,
		}
	}

	ruleMap[gitDir] = literalOrExpr{
		lit:  gitDir,
		expr: nil,
	}

	return GitIgnore{
		ruleMap:             ruleMap,
		pathToGitignoreFile: path,
	}
}

func NewFromFilePath(path string) (GitIgnore, error) {
	p := filepath.Join(path, gitignore)
	b, err := os.ReadFile(p)
	if err != nil {
		return GitIgnore{}, err
	}

	return NewGitIgnore(string(b), path), nil
}

func (gi *GitIgnore) PathExcluded(path string) bool {
	if strings.Contains(path, gitDir) {
		return true
	}

	path = gi.cutPrefix(path)

	// Maybe its just a literal so we do not have to iterate over everything.
	// Not really necessary but meh.
	loe, ok := gi.ruleMap[path]
	if ok {
		return loe.matchOrCompare(path)
	}

	for _, loe := range gi.ruleMap {
		if loe.matchOrCompare(path) {
			return true
		}
	}

	return false
}

func (gi *GitIgnore) cutPrefix(path string) string {
	af, _ := strings.CutPrefix(path, gi.pathToGitignoreFile)
	return af
}
