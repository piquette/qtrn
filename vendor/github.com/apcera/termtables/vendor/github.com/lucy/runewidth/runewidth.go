//go:generate go run gen.go override.go

// Package runewidth measures the number of cells taken up by characters on a
// terminal.
//
// RuneWidth is an implementation of wcwidth, as defined in the Single UNIX
// specification. StringWidth likewise implements wcswidth.
//
// See https://www.cl.cam.ac.uk/~mgk25/ucs/wcwidth.c for further information.
// This package is based on codegen from golang.org/x/text and the python
// implementation at https://github.com/jquast/wcwidth. It produces equivalent
// results to the python implementation for all valid inputs as of 2015-11-29.
package runewidth

var trie = newWidthTrie(0)

// RuneWidth returns the width of a character.
//
// Returns 0 for zero-width characters and characters that have no effect on
// the terminal (like NULL), -1 for control codes, and 1 or 2 for printable
// characters.
//
// Characters that return -1:
//   - C0 control codes except NULL (U+0001 - U+001F)
//   - C1 control codes (U+007F - U+009F)
//
// Characters that return 0:
//   - Unicode category Mn (Mark, nonspacing)
//   - Unicode category Me (Mark, spacing combining)
//   - NULL (U+0000)
//   - ZERO WIDTH SPACE - RIGHT-TO-LEFT MARK (U+200B - U+200F)
//   - LINE SEPARATOR - PARAGRAPH SEPERATOR (U+2028 - U+2029)
//   - LEFT-TO-RIGHT EMBEDDING - RIGHT-TO-LEFT OVERRIDE (U+202A - U+202E)
//   - WORD JOINER - INVISIBLE SEPARATOR (U+2060 - U+2063)
//   - COMBINING GRAPHEME JOINER (U+034F)
//
// Characters that return 1:
//   - SOFT HYPHEN (U+00AD)
//   - All other characters
//
// Characters that return 2:
//   - East Asian Fullwidth (F) and East Asian Wide (W) characters as defined
//     in Unicode Standard Annex #11 (http://www.unicode.org/reports/tr11/)
func RuneWidth(r rune) int {
	v := trie.lookupStringUnsafe(string(r))
	if v == 3 {
		// 3 represents -1
		return -1
	}
	// 0, 1 and 2 represent 1 and 2, and 0
	return (int(v) + 1) % 3
}

// StringWidth returns the width of a string.
//
// Returns -1 if a character RuneWidth would return -1 for is found.
//
// An incomplete encoding at the end of the string counts as 0 cells. Other
// invalid encodings count as 1 cell.
func StringWidth(s string) int {
	w := 0
	i := 0
	for i < len(s) {
		v, sz := trie.lookupString(s[i:])
		if sz == 0 {
			// incomplete code point
			// TODO: decide what to do with this
			return w
		}
		if v == 3 {
			return -1
		}
		w += (int(v) + 1) % 3
		i += sz
	}
	return w
}

// LookupString returns the terminal width of the first UTF-8 encoding in s and
// its length in bytes. The length will be 0 if s does not have enough bytes to
// complete the encoding. len(s) must be greater than 0.
func LookupString(s string) (int, int) {
	v, sz := trie.lookupString(s)
	if v == 3 {
		return -1, sz
	}
	return (int(v) + 1) % 3, sz
}

// LookupBytes returns the terminal width of the first UTF-8 encoding in s and
// its length in bytes. The length will be 0 if s does not have enough bytes to
// complete the encoding. len(s) must be greater than 0.
func LookupBytes(s []byte) (int, int) {
	v, sz := trie.lookup(s)
	if v == 3 {
		return -1, sz
	}
	return (int(v) + 1) % 3, sz
}
