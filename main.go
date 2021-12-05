package main

import (
	"fmt"
	file "jftp/filemanager"
	"log"
)

func main() {
	fmt.Print("File transfer by Jota\n\n")
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
	fileStats, err := file.GetFileStats("whom.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Name: %v\nSize: %v\nPermissions: %v\n", fileStats.Name(), fileStats.Size(), fileStats.Mode())
	fmt.Printf("Content: %v\n", fileContent)
}
