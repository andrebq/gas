#GAS

GAS (Gopath ASsets) is a lib to load resources (images, text, videos, templates, etc...) using GOPATH.

#Installing

	go get github.com/andrebq/gas

#Using

	file, err := gas.Open("import/path/to/file.png")
	// This is similar to os.Open but it search the GOPATH variable for possible locations
	// the first one to match is used
	//
	// Same idea of "GOPATH" variables but for resources and not only go packages.

#License

Released under MIT License, see LICENSE for more information
