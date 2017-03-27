package runewidth

// TODO: test malformed strings

import (
	"testing"
)

func BenchmarkSlow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = RuneWidth('コ')
	}
}

type StringTest struct {
	desc  string
	input string
	each  []int
	full  int
}

func runStringTest(t *testing.T, s StringTest) {
	each := []int{}
	for _, r := range s.input {
		each = append(each, RuneWidth(r))
	}
	full := StringWidth(s.input)
	if s.full != full {
		t.Errorf("StringWidth(%q) = %d; want %d", s.input, full, s.full)
	}
	for i := range each {
		explen := s.each[i]
		outlen := each[i]
		if explen != outlen {
			t.Errorf("RuneWidth(%q) = %v; want %v", s.input, each, s.each)
			return
		}
	}
}

var stringTests = []StringTest{
	// Tests nabbed from https://github.com/jquast/wcwidth
	// TODO: add more tests

	// simple phrase
	{input: "コンニチハ, セカイ!", each: []int{2, 2, 2, 2, 2, 1, 1, 2, 2, 2, 1}, full: 19},

	// 0 is 0 width
	{input: "abc\x00def", each: []int{1, 1, 1, 0, 1, 1, 1}, full: 6},

	// CSI string is -1 width
	{input: "\x1b[0m", each: []int{-1, 1, 1, 1}, full: -1},

	// combining char is 0 width
	{input: "--\u05bf--", each: []int{1, 1, 0, 1, 1}, full: 4},

	// cafe + COMBINING ACUTE ACCENT is café, width 4
	{input: "cafe\u0301", each: []int{1, 1, 1, 1, 0}, full: 4},

	// CYRILLIC CAPITAL LETTER A + COMBINING CYRILLIC HUNDRED THOUSANDS SIGN
	// is А҈, width 1.
	{input: "\u0410\u0488", each: []int{1, 0}, full: 1},

	// Balinese kapal (ship) is ᬓᬨᬮ᭄ , width 44
	{input: "\u1B13\u1B28\u1B2E\u1B44", each: []int{1, 1, 1, 1}, full: 4},
}

func TestStrings(t *testing.T) {
	for _, s := range stringTests {
		runStringTest(t, s)
	}
}
