package filemanager

import "os"

func DeleteFile(filename string) error {
	err := os.Remove("test.txt")
	return err
}

func CopyFile(filenameToCopy string, newFilename string) error {
	content, err := ReadFile(filenameToCopy)
	if err != nil {
		return err
	}
	return CreateAndWriteFile(newFilename, string(content))
}

func RenameFile(filename string, newFilename string) error {
	return os.Rename(filename, newFilename)
}
