package version

// Package is the overall, canonical project import path under which the
var Package = "github.com/piquette/qtrn"

// Version indicates which version of the binary is running. This is set to
// the latest release tag by hand, always suffixed by "+unknown". During
// build, it will be replaced by the actual version. The value here will be
// used if the registry is run after a go get based install.
var Version = "master"
