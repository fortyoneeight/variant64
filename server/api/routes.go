package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/variant64/server/player"
	"github.com/variant64/server/room"
)

const InvalidBodyResponse = "{\"error\":\"invalid request body\"}"

func AttachRESTRoutes(r *mux.Router) {
	r.HandleFunc("/api/player", handlePostPlayer).Methods("POST")
	r.HandleFunc("/api/players/{id}", handleGetPlayer).Methods("GET")

	r.HandleFunc("/api/room", handlePostRoom).Methods("POST")
	r.HandleFunc("/api/rooms", handleGetRooms).Methods("GET")

	r.HandleFunc("/api/room/{room_name}/join", handlePostRoomJoin).Methods("POST")
	r.HandleFunc("/api/room/{room_name}/leave", handlePostRoomLeave).Methods("POST")
	r.HandleFunc("/api/room/{room_name}/start", handlePostRoomStart).Methods("POST")
}

type RequestPostPlayer struct {
	DisplayName string `json:"display_name"`
}

func handlePostPlayer(w http.ResponseWriter, req *http.Request) {
	body := RequestPostPlayer{}
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil || body.DisplayName == "" {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	playerStore := player.GetPlayerStore()

	newPlayer := player.NewPlayer(body.DisplayName)
	playerStore.Lock()
	playerStore.Store(newPlayer)
	playerStore.Unlock()

	serialized, _ := json.Marshal(newPlayer)

	w.Write(serialized)
}

func handleGetPlayer(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	playerStore := player.GetPlayerStore()

	playerStore.Lock()
	player := playerStore.GetByID(id)
	playerStore.Unlock()

	serialized, _ := json.Marshal(player)

	w.Write(serialized)
}


type RequestPostRoom struct {
	RoomName string `json:"room_name"`
}

func handlePostRoom(w http.ResponseWriter, req *http.Request) {
	body := RequestPostRoom{}
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil || body.RoomName == "" {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	roomStore := room.GetRoomStore()

	newRoom := room.NewRoom(body.RoomName)
	roomStore.Lock()
	roomStore.Store(newRoom)
	roomStore.Unlock()


	serialized, _ := json.Marshal(newRoom)

	w.Write(serialized)
}

func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	roomStore := room.GetRoomStore()

	roomStore.Lock()
	allRooms := roomStore.GetAll()
	roomStore.Unlock()

	serialized, _ := json.Marshal(allRooms)

	w.Write(serialized)
}

type RequestPostRoomPlayer struct {
	PlayerID uuid.UUID `json:"player_id"`
}

func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	body := RequestPostRoomPlayer{}

	id, err := uuid.Parse(vars["room_name"])
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	err = json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	playerStore := player.GetPlayerStore()
	playerStore.Lock()
	players := playerStore.GetByID(body.PlayerID)
	if len(players) == 0 {
		playerStore.Unlock()
		w.Write([]byte("{\"error\":\"player not found\"}"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	player := *players[0]
	playerStore.Unlock()

	roomStore := room.GetRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()
	rooms := roomStore.GetByID(id)
	if len(rooms) == 0 {
		w.Write([]byte("{\"error\":\"room not found\"}"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	room := rooms[0]
	room.AddPlayer(player)

	serialized, _ := json.Marshal(room)

	w.Write(serialized)
}

func handlePostRoomLeave(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	body := RequestPostRoomPlayer{}

	id, err := uuid.Parse(vars["room_name"])
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	err = json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	roomStore := room.GetRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()
	rooms := roomStore.GetByID(id)
	if len(rooms) == 0 {
		w.Write([]byte("{\"error\":\"room not found\"}"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	room := rooms[0]
	err = room.RemovePlayer(body.PlayerID)
	if err != nil {
		w.Write([]byte("{\"error\":\"player not in room\"}"))
		return
	}

	serialized, _ := json.Marshal(room)

	w.Write(serialized)
}

func handlePostRoomStart(w http.ResponseWriter, req *http.Request) {
	response := "{}"
	serialized, _ := json.Marshal(response)

	w.Write(serialized)
}
