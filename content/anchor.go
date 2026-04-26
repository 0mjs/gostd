package content

import (
	"strconv"
	"strings"
	"unicode"
)

// Slugify produces a URL-safe anchor from arbitrary text. Used by both the
// renderer (main.go) and the search-index builder so anchors stay in sync.
func Slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var b strings.Builder
	lastDash := false
	for _, r := range s {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			b.WriteRune(r)
			lastDash = false
		case !lastDash:
			b.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

// UniqueAnchor returns a slug guaranteed unique against the supplied counter
// map. Mutates the map. Mirrors the renderer behavior in main.go.
func UniqueAnchor(used map[string]int, raw string) string {
	base := Slugify(raw)
	if base == "" {
		base = "section"
	}
	if used[base] == 0 {
		used[base] = 1
		return base
	}
	used[base]++
	return base + "-" + strconv.Itoa(used[base])
}
