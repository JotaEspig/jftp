package filemanager

import (
	"io/ioutil"
)

func WriteFile(filename string, textoToWrite string) error {
	return ioutil.WriteFile(filename, []byte(textoToWrite+"\n"), 0666)
}
