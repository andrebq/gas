// +build go1.1

package gas

import (
	"go/build"
	"path/filepath"
)

func gopathDirs() []string {
	return filepath.SplitList(build.Default.GOPATH)
}
