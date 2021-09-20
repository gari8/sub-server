package file_manager

import (
	"io/ioutil"
	"os"
)

type FileManager struct {
	FileName string
}

func NewFileManager(filePath string) FileManager {
	return FileManager{
		FileName: filePath,
	}
}

func (fm FileManager) Create(content []byte) error {
	file, err := os.Create(fm.FileName)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func (fm FileManager) Read() ([]byte, error) {
	return ioutil.ReadFile(fm.FileName)
}
