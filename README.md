#GAS

GAS (Gopath ASsets) is a lib to load resources (images, text, videos, templates, etc...) using GOPATH.

#Installing

	go get github.com/andrebq/gas

#Using

	file, err := gas.Open("import/path/to/file.png")
	// This is similar to os.Open but it search the GOPATH variable for possible locations
	// the first one to match is used
	//
	// The matched file will be: $GOPATH/src/import/path/to/file.png
	//
	// Same idea of "GOPATH" variables but for resources and not only go packages.
	//
	// If the file isn't found on gopath, this will search the current folder for the location.

The behavior is well defined:

First it search the gopath (prepending the src base folder), if the file isn't found then it will search the current folder.

#License

Released under MIT License, see LICENSE for more information
