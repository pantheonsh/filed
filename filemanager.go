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
	var scanpath = fm.CurrPath
	var rfiles = []File{}
	files, err := ioutil.ReadDir(scanpath)

	if err != nil {
		return nil
	}

	for _, elem := range files {
		if elem.IsDir() {
			continue
		}
		file := File{
			Name: elem.Name(),
			Path: scanpath,
			Size: elem.Size(),
		}
		rfiles = append(rfiles, file)
	}

	return rfiles
}
