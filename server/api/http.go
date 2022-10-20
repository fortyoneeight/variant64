package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/variant64/server/entity"
)

const InvalidBodyResponse = "{\"error\":\"invalid request body\"}"

func handlePostPlayer(w http.ResponseWriter, req *http.Request) {
	request := &entity.RequestNewPlayer{}
	handleWriteEntity[entity.Player](w, req, request)
}

func handleGetPlayerByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	request := &entity.RequestGetPlayer{ID: id}
	handleReadEntity[entity.Player](w, req, request)
}

func handlePostRoom(w http.ResponseWriter, req *http.Request) {
	request := &entity.RequestNewRoom{}
	handleWriteEntity[entity.Room](w, req, request)
}

func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	request := &entity.RequestGetRooms{}
	handleReadEntities[entity.Room](w, req, request)
}

func handleGetRoomByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	request := &entity.RequestGetRoom{ID: id}
	handleReadEntity[entity.Room](w, req, request)
}

func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	request := &entity.RequestRoomAddPlayer{RoomID: id}
	handleWriteEntity[entity.Room](w, req, request)
}

func handlePostRoomLeave(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	request := &entity.RequestRoomRemovePlayer{RoomID: id}
	handleWriteEntity[entity.Room](w, req, request)
}

func handlePostRoomStart(w http.ResponseWriter, req *http.Request) {
	response := "{}"
	serialized, _ := json.Marshal(response)

	w.Write(serialized)
}
