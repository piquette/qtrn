// Generates the zero and double width RangeTables

// +build ignore

package main

import (
	"bytes"
	"log"
	"strings"

	"golang.org/x/text/internal/gen"
	"golang.org/x/text/internal/triegen"
	"golang.org/x/text/internal/ucd"
)

const (
	widthOne  = 0
	widthTwo  = 1
	widthZero = 2
	widthNil  = 3
)

type override struct {
	from, to rune
	width    int
}

func encodeWidth(w int) uint64 {
	switch w {
	case -1:
		return widthNil
	case 0:
		return widthZero
	case 1:
		return widthOne
	case 2:
		return widthTwo
	}
	panic("invalid width")
}

func contains(s string, ps ...string) bool {
	for _, p := range ps {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}

func main() {
	t := triegen.NewTrie("width")

	// wide is the base
	parse("EastAsianWidth.txt", func(p *ucd.Parser) {
		if contains(p.String(1), "W", "F") {
			t.Insert(p.Rune(0), widthTwo)
		}
	})

	// zero overrides wide
	parse("extracted/DerivedGeneralCategory.txt", func(p *ucd.Parser) {
		cat := p.String(1)
		if cat == "Me" || cat == "Mn" {
			t.Insert(p.Rune(0), widthZero)
		}
	})

	// misc overrides
	for _, v := range overrides {
		for r := v.from; r <= v.to; r++ {
			t.Insert(r, encodeWidth(v.width))
		}
	}

	w := &bytes.Buffer{}
	gen.WriteUnicodeVersion(w)
	t.Gen(w)
	gen.WriteGoFile("tables.go", "runewidth", w.Bytes())
}

func parse(path string, f func(p *ucd.Parser)) {
	r := gen.OpenUCDFile(path)
	defer r.Close()
	p := ucd.New(r)
	for p.Next() {
		f(p)
	}
	if err := p.Err(); err != nil {
		log.Fatal(err)
	}
}
