package gas

import (
	"io"
	"io/ioutil"
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

// Return the absolute filepath for the requested resource or return an error if not found
func Abs(file string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()
	return fs.Abs(file, true)
}

// MustAbs ensure that the given file is present in the system, if the file can't
// be found, this will call panic giving the reason
func MustAbs(file string) string {
	ret, err := Abs(file)
	if err != nil {
		panic(ret)
	}
	return ret
}

// ReadFile return the contents of the file at the given gopath
func ReadFile(file string) ([]byte, error) {
	rc, err := Open(file)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}
