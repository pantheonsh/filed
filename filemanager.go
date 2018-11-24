package main

import (
	"io/ioutil"
)

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

func (fm FileManager) getDirectoryFiles() []File {
	files, err := ioutil.ReadDir(fm.CurrPath)
	var rfiles = []File{}

	if err != nil {
		return nil
	}

	for _, elem := range files {
		if elem.IsDir() {
			continue
		}
		file := File{
			Name: "oi",
			Path: "d",
			Size: elem.Size(),
		}
		rfiles = append(rfiles, file)
	}

	return rfiles
}
