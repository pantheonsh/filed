package main

import "io/ioutil"

// FileManager type
type FileManager struct {
	CurrPath string
}

// File type
type File struct {
	Name string
	Path string
	Size int64
}

func (fm FileManager) getDirectoryListing() []File {
	files, err := ioutil.ReadDir(fm.CurrPath)
	var rfiles = []File{}

	if err != nil {
		return nil
	}

	for _, elem := range files {
		file := File{
			Name: "oi",
			Path: "d",
			Size: elem.Size(),
		}
		_ = append(rfiles, file)
	}

	return rfiles
}
