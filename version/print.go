package version

import (
	"fmt"
	"io"
	"os"
)

// FprintVersion outputs the version string to the writer.
func FprintVersion(w io.Writer) {
	fmt.Fprintln(w, os.Args[0], Package, Version)
}

// PrintVersion outputs the version information, from Fprint, to stdout.
func PrintVersion() {
	FprintVersion(os.Stdout)
}
