package runewidth

import (
	"testing"

	lucy "github.com/lucy/runewidth"
	mattn "github.com/mattn/go-runewidth"
)

func BenchmarkEasyRune(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = lucy.RuneWidth('a')
	}
}

func BenchmarkEasyString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = lucy.StringWidth("abcdefgkljjsfkjn")
	}
}

func Benchmark1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = lucy.RuneWidth('コ')
	}
}

func Benchmark2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = lucy.StringWidth("■㈱の世界①")
	}
}

func Benchmark3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = lucy.StringWidth("あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおおおおお")
	}
}

const long = `ああいうえおあいうえおえおおおおおおおおおおおおおおおあいう` +
	`えおあいうえおえおおおおおおおおおおおおおおおあいうえおあいうえお` +
	`えおおおおおおおおおおおおおおおあいうえおあいうえおえおおおおおお` +
	`おおおおおおおおおあいうえおあいうえおえおおおおおおおおおおおおお` +
	`おおあいうえおあいうえおえおおおおおおおおおおおおおおおあいうえお` +
	`あいうえおえおおおおおおおおおおおおおおおあいうえおあいうえおえお` +
	`おおおおおおおおおおおおおおあいうえおあいうえおえおおおおおおおお` +
	`おおおおおおおあいうえおあいうえおえおおおおおおおおおおおおおおお` +
	`あいうえおあいうえおえおおおおおおおおおおおおおおおあいうえおあい` +
	`うえおえおおおおおおおおおおおおおおおあいうえおあいうえおえおおお` +
	`おおおおおおおおおおおおあいうえおあいうえおえおおおおおおおおおお` +
	`おおおおおあいうえおあいうえおえおおおおおおおおおおおおおおおあい` +
	`うえおあいうえおえおおおおおおおおおおおおおおおあいうえおあいうえ` +
	`おえおおおおおおおおおおおおおおおあいうえおあいうえおえおおおおお` +
	`おおおおおおおおおおあいうえおあいうえおえおおおおおおおおおおおお` +
	`おおおあいうえおあいうえおえおおおおおおおおおおおおおおおあいうえ` +
	`おあいうえおえおおおおおおおおおおおおおおおあいうえおあいうえおえ` +
	`おおおおおおおおおおおおおおおあいうえおあいうえおえおおおおおおお` +
	`おおおおおおおおあいうえおあいうえおえおおおおおおおおおおおおおお` +
	`おあいうえおあいうえおえおおおおおおおおおおおおおおおあいうえおあ` +
	`いうえおえおおおおおおおおおおおおおおおあいうえおあいうえおえおお` +
	`おおおおおおおおおおおおおあいうえおあいうえおえおおおおおおおおお` +
	`おおおおおおあいうえおあいうえおえおおおおおおおおおおおおおおおあ` +
	`いうえおあいうえおえおおおおおおおおおおおおおおおあいうえおあいう` +
	`えおえおおおおおおおおおおおおおおおいうえおあいうえおえおおおおお` +
	`おおおおおおおおおおおおおおおおおおおおおおおおお`

func Benchmark4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = lucy.StringWidth(long)
	}
}

func BenchmarkOtherlibEasyRune(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mattn.RuneWidth('a')
	}
}

func BenchmarkOtherlibEasyString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mattn.StringWidth("abcdefgkljjsfkjn")
	}
}

func BenchmarkOtherlib1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mattn.RuneWidth('コ')
	}
}

func BenchmarkOtherlib2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mattn.StringWidth("■㈱の世界①")
	}

}

func BenchmarkOtherlib3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mattn.StringWidth("あいうえおあいうえおえおおおおおおおおおおおおおおおおおおおおおおおおおおおおおお")
	}
}

func BenchmarkOtherlib4(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mattn.StringWidth(long)
	}
}
