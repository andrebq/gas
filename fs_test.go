package gas

import (
	"os"
	"testing"
)

func TestAbs(t *testing.T) {
	fs := UnitedFS()
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

	abs, err = fs.Abs("fs.go", false)
	if err != nil {
		t.Fatalf("Should have found the fs.go file. But got: %v", err)
	}

	if abs == "" {
		t.Fatalf("abs should have a valid path inside of it")
	}

	info, err = os.Stat(abs)
	if err != nil {
		t.Fatalf("os.Stat returned a error: %v", err)
	}

	if info.IsDir() {
		t.Fatalf("abs is pointing to a directory: (%v)", abs)
	}
}

func TestAbsDir(t *testing.T) {
	fs := UnitedFS()
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

	abs, err = fs.Abs("../gas", true)
	if err != nil {
		t.Fatalf("Shoudl have found the ../gas directory. But got: %v", err)
	}

	if abs == "" {
		t.Fatalf("abs should have a valid path inside of it")
	}

	info, err = os.Stat(abs)
	if err != nil {
		t.Fatalf("os.Stat returned a error: %v", err)
	}
}

func TestOpen(t *testing.T) {
	fs := UnitedFS()
	file, err := fs.Open("github.com/andrebq/gas/fs.go")
	if err != nil {
		t.Fatalf("Unable to open fs.go file. %v", err)
	}
	file.Close()

	file, err = fs.Open("fs.go") // should find the file using the "." path
	if err != nil {
		t.Fatalf("Unable to open fs.go (using local folder). %v", err)
	}
	file.Close()
}

func TestFromDirs(t *testing.T) {
	ufs := UnitedFS()
	gasDir, err := ufs.Abs("github.com/gas", true)

	dirfs := FromDirs([]string{gasDir})

	// since I added github.com/gas as a directory using
	// unitedfs
	// the dirfs should be able to find the fs.go file
	// without any prefix
	file, err := dirfs.Open("fs.go")
	if err != nil {
		t.Fatalf("should have found the fs.go file")
	}
	defer file.Close()
}
