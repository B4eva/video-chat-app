package main

import (
	// "fmt"
	"log"
	"net/http"
	"video-chat-app/server"
)

func main() {
	http.HandleFunc("/create", server.CreatRoomRequestHandler)
	http.HandleFunc("/join", server.JoinRoomRequestHandler)

	log.Println(("Starting server on Port 8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
