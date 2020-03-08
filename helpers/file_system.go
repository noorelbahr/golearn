package helpers

import "net/http"

type MyFS struct {
	http.Dir
}

/**
 * Custom File System Dir Open function
 */
func (m MyFS) Open(name string) (result http.File, err error) {
	// Open file
	file, err := m.Dir.Open(name)
	if err != nil {
		return
	}

	// Get file info
	fileState, err := file.Stat()
	if err != nil {
		return
	}

	// Check whether the requested file is a directory or not
	if fileState.IsDir() {
		// If so, return a response 404 (cause "does-not-exist" is doesn't exist in our server)
		return m.Dir.Open("does-not-exist")
	}

	return file, nil
}