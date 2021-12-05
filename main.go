package main

import (
	"fmt"
	file "jftp/filemanager"
	"log"
)

func main() {
	fmt.Println("File transfer by Jota")
	fileContent, err := file.ReadFile("whom.txt")
	if err != nil {
		err = file.CreateAndWriteFile("whom.txt", "Vai tomar no cu seu corno do caralho")
		if err != nil {
			log.Fatalln(err)
		}
		fileContent, err = file.ReadFile("whom.txt")
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println(fileContent)
}
