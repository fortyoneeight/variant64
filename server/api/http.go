package api

import (
	"net/http"

	"github.com/variant64/server/entity"
)

func handlePostPlayer(w http.ResponseWriter, req *http.Request) {
	handleNewEntity[entity.Player](w, req, &entity.RequestNewPlayer{})
}

func handleGetPlayerByID(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	handleGetEntity[entity.Player](w, req, &entity.RequestGetPlayer{ID: id})
}

func handlePostRoom(w http.ResponseWriter, req *http.Request) {
	handleNewEntity[entity.Room](w, req, &entity.RequestNewRoom{})
}

func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	requestGetRooms := &entity.RequestGetRooms{}
	rooms := requestHandler.HandleGetRooms(requestGetRooms)

	writeEntityListResponse(w, rooms)
}

func handleGetRoomByID(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	handleGetEntity[entity.Room](w, req, &entity.RequestGetRoom{ID: id})
}

func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	handleActionRouteByID[entity.Room](w, req, &entity.RequestJoinRoom{RoomID: id})
}

func handlePostRoomLeave(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	handleActionRouteByID[entity.Room](w, req, &entity.RequestLeaveRoom{RoomID: id})
}

func handlePostRoomStart(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	handleActionRouteByID[entity.Room](w, req, &entity.RequestStartGame{RoomID: id})
}
