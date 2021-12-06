package main

import (
	"jftp/socket"
	"log"
)

func main() {
	log.Println("JFTP listening any on port 5001")
	server := socket.CreateServer("0.0.0.0", 5001)
	err := server.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
