package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const indentSpaces = 2

// parsePaths converts the indented tree printed by `keepassxc-cli ls -R` into
// full entry paths. Group lines end in "/", and each nesting level is indented
// by indentSpaces.
func parsePaths(scanner *bufio.Scanner) []string {
	paths := []string{}
	groups := []string{} // group name per level; index == level

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimLeft(line, " ")
		if trimmed == "" {
			continue
		}
		level := (len(line) - len(trimmed)) / indentSpaces

		// Dedent: drop every group deeper than this line's level.
		if level < len(groups) {
			groups = groups[:level]
		}

		if strings.HasSuffix(trimmed, "/") {
			groups = append(groups, trimmed)
			continue
		}
		paths = append(paths, strings.Join(groups, "")+trimmed)
	}
	return paths
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	paths := parsePaths(scanner)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		os.Exit(1)
	}
	for _, path := range paths {
		fmt.Println(path)
	}
}
