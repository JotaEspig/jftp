package filemanager

import "os"

func CreateEmptyFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	file.Close()
	err = os.Chmod(filename, 0666)
	return err
}

func CreateAndWriteFile(filename string, textToWrite string) error {
	err := CreateEmptyFile(filename)
	if err != nil {
		return err
	}
	err = WriteFile(filename, textToWrite)
	return err
}
