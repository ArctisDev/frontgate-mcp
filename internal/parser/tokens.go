package parser

import (
	"regexp"
	"sort"
	"strings"
)

var tokenPattern = regexp.MustCompile(`--([a-zA-Z0-9-_]+)\s*:\s*([^;]+);`)

func ExtractCSSVariables(body string) map[string]string {
	matches := tokenPattern.FindAllStringSubmatch(body, -1)
	out := make(map[string]string, len(matches))
	for _, m := range matches {
		out[m[1]] = strings.TrimSpace(m[2])
	}
	return out
}

var spacingClassPattern = regexp.MustCompile(`\b(?:p|px|py|pt|pr|pb|pl|m|mx|my|mt|mr|mb|ml|gap|space-x|space-y)-\[[^\]]+\]|\b(?:p|px|py|pt|pr|pb|pl|m|mx|my|mt|mr|mb|ml|gap|space-x|space-y)-[0-9]+(?:\.[05])?\b`)

func ExtractSpacingClasses(body string) []string {
	matches := spacingClassPattern.FindAllString(body, -1)
	sort.Strings(matches)
	return matches
}
