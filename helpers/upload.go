package helpers

import (
	"io/ioutil"
	"mime/multipart"
)

func Upload(file multipart.File, handler *multipart.FileHeader, dir string) (string, error) {
	// Make temp file
	tempFile, err := ioutil.TempFile(dir, "*-" + handler.Filename)
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// Read the file
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Write file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return "", err
	}

	return tempFile.Name(), nil
}
