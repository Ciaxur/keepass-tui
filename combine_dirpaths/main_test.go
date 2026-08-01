package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestParsePaths(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  []string
	}{
		{"flat entries", "one\ntwo\n", []string{"one", "two"}},
		{
			"dedent after nested group",
			"Root/\n  a\n  Sub/\n    b\n  c\nOther/\n  d\ntop\n",
			[]string{"Root/a", "Root/Sub/b", "Root/c", "Other/d", "top"},
		},
		{"empty group", "Empty/\nafter\n", []string{"after"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := parsePaths(bufio.NewScanner(strings.NewReader(tc.input)))
			if len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("path %d: got %q, want %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}
