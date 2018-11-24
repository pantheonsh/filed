package main

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
	return nil
}
