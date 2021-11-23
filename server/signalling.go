package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "github.com/gorilla/websocket"
)

var AllRooms RoomMap

//CreateRoomRequestHandler Create a Room and return roomID
func CreatRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	roomID := AllRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}
	log.Println(AllRooms.Map)

	json.NewEncoder(w).Encode(resp{RoomID: roomID})

}

//JoinRoomRequestHandler will join the client in a particular room
func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}
