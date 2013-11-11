// +build !go1.1

package gas

import (
	"go/build"
	"os"
	"strings"
)

func gopathDirs() []string {
	return strings.Split(build.Default.GOPATH, string(os.PathListSeparator))
}
