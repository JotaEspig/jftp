package filemanager

import "os"

func DeleteFile(filename string) error {
	return os.Remove("test.txt")
}

func CopyFile(filenameToCopy string, newFilename string) error {
	content, err := ReadFile(filenameToCopy)
	if err != nil {
		return err
	}
	return CreateAndWriteFile(newFilename, string(content))
}

func MoveFile(filename string, newFilename string) error {
	err := CopyFile(filename, newFilename)
	if err != nil {
		return err
	}
	return DeleteFile(filename)
}

func RenameFile(oldFilename string, newFilename string) error {
	return os.Rename(oldFilename, newFilename)
}

func GetFileStats(filename string) (os.FileInfo, error) {
	stats, err := os.Stat(filename)
	return stats, err
}
