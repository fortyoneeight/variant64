package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AttachRoutes(r *mux.Router) {
	r.HandleFunc("/api/room", handlePostRoom).Methods("POST")
	r.HandleFunc("/api/rooms", handleGetRooms).Methods("GET")
	r.HandleFunc("/api/room/{room_name}/join", handlePostRoomJoin).Methods("POST")
	r.HandleFunc("/api/room/{room_name}/start", handlePostRoomStart).Methods("POST")
}

func handlePostRoom(w http.ResponseWriter, req *http.Request) {
	response := "{\"room_name\":\"Room\",\"players_total\":2,\"players\":[]}}"
	serialized, _ := json.Marshal(response)

	w.Write(serialized)
}

func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	response := "{\"rooms\":[{\"room_name\":\"Room\",\"players_total\":2,\"players\":[]}}]"
	serialized, _ := json.Marshal(response)

	w.Write(serialized)
}

func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	body := make(map[string]string)
	_ = json.NewDecoder(req.Body).Decode(&body)

	fmt.Println(vars)
	response := fmt.Sprintf("{\"room_name\":\"%s\",\"players_total\":2,\"players\":[\"%s\"]}", vars["room_name"], body["player_name"])
	serialized, _ := json.Marshal(response)

	w.Write(serialized)
}

func handlePostRoomStart(w http.ResponseWriter, req *http.Request) {
	response := "{}"
	serialized, _ := json.Marshal(response)

	w.Write(serialized)
}
