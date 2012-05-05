package gas

import (
	"os"
	"testing"
)

func TestAbs(t *testing.T) {
	fs := GopathFS()
	abs, err := fs.Abs("github.com/andrebq/gas/fs.go", false)
	if err != nil {
		t.Fatalf("Should have found the fs.go file. But got: %v", err)
	}

	if abs == "" {
		t.Fatalf("abs should have a valid path inside of it")
	}

	info, err := os.Stat(abs)
	if err != nil {
		t.Fatalf("os.Stat returned a error: %v", err)
	}

	if info.IsDir() {
		t.Fatalf("abs is pointing to a directory: (%v)", abs)
	}
}

func TestAbsDir(t *testing.T) {
	fs := GopathFS()
	abs, err := fs.Abs("github.com/andrebq/gas/", true)
	if err != nil {
		t.Fatalf("Should have found the gas directory. But got: %v", err)
	}

	if abs == "" {
		t.Fatalf("abs should have a valid path inside of it")
	}

	info, err := os.Stat(abs)
	if err != nil {
		t.Fatalf("os.Stat returned a error: %v", err)
	}

	if !info.IsDir() {
		t.Fatalf("abs isn't pointing to a directory: (%v)", abs)
	}
}

func TestOpen(t *testing.T) {
	fs := GopathFS()
	file, err := fs.Open("github.com/andrebq/gas/fs.go")
	if err != nil {
		t.Fatalf("Unable to open fs.go file. %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Fatalf("Error while closing the file. %v", err)
	}
}
