package filemanager

import (
	"io/ioutil"
)

func WriteFile(filename string, textoToWrite string) error {
	err := ioutil.WriteFile(filename, []byte(textoToWrite+"\n"), 0666)
	return err

}
