package gas

import (
	"io"
	"sync"
)

var (
	fs   *FS
	lock sync.RWMutex
)

// Initialize the code using the default (UnitedFS)
func init() {
	fs = UnitedFS()
}

// Refresh the internal FS to reflect possible changes in the UnitedFS
func Refresh() {
	lock.Lock()
	defer lock.Unlock()
	fs = UnitedFS()
}

// Open the file for reading or returns an error
//
// For more information, see the FS type
func Open(file string) (io.ReadCloser, error) {
	lock.RLock()
	defer lock.RUnlock()
	return fs.Open(file)
}

// Return the absolut filepath for the requested resource or return an error if not found
func Abs(file string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	return fs.Abs(file, true)
}
