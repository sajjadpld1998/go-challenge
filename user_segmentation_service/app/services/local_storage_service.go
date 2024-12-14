package services

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"skeleton/config"
)

func (localStorage LocalStorage) Store(file *multipart.FileHeader, Path string) (err error) {
	directoryPath := config.GetConfig().FileSystem.Disks.Local.Root + Path

	directoryByFilePath := directoryPath + file.Filename

	storeFile, err := file.Open()
	if err != nil {
		return
	}
	defer storeFile.Close()

	err = os.MkdirAll(directoryPath, 0755)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(storeFile)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(directoryByFilePath, data, 0777)
	if err != nil {
		return
	}

	return
}
