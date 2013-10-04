package gas

import (
	"os"
	"strings"
)

func gopathDirs() []string {
	return strings.Split(os.Getenv("GOPATH"), ":")
}
