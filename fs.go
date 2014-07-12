package gas

import (
	"io"
	"os"
	"path"
	"path/filepath"
)

type FS struct {
	searchPath []string
}

// Used to indicat that the file wasn't found in any possible location
type NotFound string

func (n NotFound) Error() string {
	return "The file " + string(n) + " wasn't found"
}

// Return true if the error represents a NotFound error
func IsNotFound(err error) bool {
	_, ok := err.(NotFound)
	return ok
}

// Find the absolute path for the required file.
//
// The returned string is OS depended. If the desired file isn't present
// in any possible location returns NotFound error.
func (fs *FS) Abs(file string, allowDir bool) (abs string, err error) {
	reqPath := filepath.FromSlash(path.Clean(file))

	for _, p := range fs.searchPath {
		abs = filepath.Join(p, reqPath)
		var stat os.FileInfo
		stat, err = os.Stat(abs)
		if !os.IsNotExist(err) {
			if !stat.IsDir() {
				return
			} else if allowDir {
				// in case the caller want's a directory
				// instead of a file
				return
			}
		}
	}
	// if reach this point
	// all possible locations were tested
	// and no match was found
	abs = ""
	err = NotFound(reqPath)
	return
}

// Open the resource for reading
func (fs *FS) Open(file string) (r io.ReadCloser, err error) {
	abs, err := fs.Abs(file, false)
	if err != nil {
		return
	}

	r, err = os.Open(abs)
	return
}

// Create a new GopathFS instance
func GopathFS() *FS {
	fs := &FS{}
	vals := gopathDirs()
	if len(vals) > 0 {
		fs.searchPath = make([]string, len(vals))
		for i, v := range vals {
			fs.searchPath[i] = filepath.Join(filepath.FromSlash(v), "src")
		}
	}
	return fs
}

// Create a FS which is the combination of the GopathFS and the current folder
// note that the current folder don't require the "src" sub-folder
func UnitedFS() *FS {
	gofs := GopathFS()
	gofs.searchPath = append(gofs.searchPath, ".")
	return gofs
}

// FromDirs return a new filesystem that searchs for the request file
// on the specified dirs
//
// This make a copy of the input array
func FromDirs(dirs []string) *FS {
	fs := &FS{}
	fs.searchPath = make([]string, len(dirs))
	copy(fs.searchPath, dirs)
	return fs
}
