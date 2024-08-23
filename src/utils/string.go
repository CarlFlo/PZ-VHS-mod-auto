package utils

import (
	"strings"
)

func ExtractStringBetweenSep(line, sepStart, sepEnd string) string {
	start := strings.Index(line, sepStart)
	if start == -1 {
		return ""
	}

	end := strings.Index(line[start+1:], sepEnd)
	if end == -1 {
		return ""
	}

	// Extract the string between the separators
	return line[start+1 : start+end+1]
}
