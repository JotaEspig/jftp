package filemanager

import "os"

func DeleteFile(filename string) error {
	err := os.Remove("test.txt")
	return err
}
