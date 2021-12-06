package main

import (
	"jftp/socket"
	"log"
	"os"
)

func main() {
	PORT := 5001
	file, err := os.OpenFile("logs.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf("JFTP Server started listening any on port %v\n", PORT)
	server := socket.CreateServer("0.0.0.0", PORT)
	err = server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
