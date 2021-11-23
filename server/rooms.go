package server

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

//Init initialises the RoomMap struct
func (r *RoomMap) Init() {

	r.Map = make(map[string][]Participant)

}

//Get will return the array of participants in the room
func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	return r.Map[roomID]

}

//CreateRoom generate a unique ID and returns it then  -> insert into the hashMap (RoomMap)
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	fmt.Println(letters)
	return string(letters)
}

func (r *RoomMap) DeleteRoom() {}
