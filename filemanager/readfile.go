package filemanager

import "os"

func ReadFile(filename string) (string, error) {
	content, err := ReadFileInBytes(filename)
	return string(content), err
}

func ReadFileInBytes(filename string) ([]byte, error) {
	var content []byte
	file, err := os.Open(filename)
	if err != nil {
		return content, err
	}
	defer file.Close()
	buff := make([]byte, 100)
	for byteRead, err := file.Read(buff); err == nil; byteRead, err = file.Read(buff) {
		if byteRead > 0 {
			content = append(content, buff[0:byteRead]...)
		}
	}
	return content, err
}
