// +build go1.1 go1.2

package gas

import (
	"go/build"
	"path/filepath"
)

func gopathDirs() []string {
	return filepath.SplitList(build.Default.GOPATH)
}
