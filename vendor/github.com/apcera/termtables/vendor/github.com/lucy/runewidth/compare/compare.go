// Compares RuneWidth with stdin (formatted as "%02x %d" % (character, width))

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lucy/runewidth"
)

func atoi(s string) int64 {
	x, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return x
}

func rsw(r rune) int {
	return runewidth.StringWidth(string(r))
}

func main() {
	var stringWidth = flag.Bool("sw", false, "test StringWidth")
	flag.Parse()
	f := runewidth.RuneWidth
	if *stringWidth {
		f = rsw
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l := s.Text()
		in := strings.Fields(l)
		if len(in) != 2 {
			fmt.Fprintf(os.Stderr, "invalid input: %q\n", l)
			os.Exit(1)
		}
		r, w := rune(atoi(in[0])), atoi(in[1])
		ow := f(r)
		if ow != int(w) {
			fmt.Fprintf(os.Stderr, "f(%q (%02x)) = %d; want %d\n", r, r, ow, w)
			//os.Exit(1)
		}
		//fmt.Fprintf(os.Stderr, "f(%q (%02x)) = %d == %d\n", r, r, ow, w)
	}
	err := s.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
