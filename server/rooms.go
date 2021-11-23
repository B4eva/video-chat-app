package server

import (
	// "fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Participant describes a single entity in the hasMap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

//RoomMap is the main hashMap [roomID string -> []Participants]
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

//CreateRoom generate a unique ID and returns it then  -> insert into the hashMap [RoomMap]
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomId := string(b)
	r.Map[roomId] = []Participant{}

	return roomId

}

// InsertIntoRoom will create a participant and add it in the hashmap
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	log.Println("inserting into Room with RoomID:", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)
}

func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}
