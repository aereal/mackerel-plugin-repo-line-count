package main

import (
	"os/exec"
	"strconv"
	"strings"
)

type MatchEntry struct {
	File  string
	Count int
}

func (e *MatchEntry) SafeMetricNameComponent() string {
	return e.File
}

type MatchResult []*MatchEntry

func Grep(path string, pattern string) MatchResult {
	var result MatchResult
	cmd := exec.Command("git", "-C", path, "grep", "-I", "--count", pattern)
	outBytes, _ := cmd.Output()
	lines := strings.Split(string(outBytes), "\n")
	for _, l := range lines {
		if l == "" {
			continue
		}
		pair := strings.SplitN(l, ":", 2)
		file := pair[0]
		count, err := strconv.Atoi(pair[1])
		if err != nil {
			continue // TODO
		}
		entry := &MatchEntry{File: file, Count: count}
		result = append(result, entry)
	}
	return result
}
