package api

import (
	"net/http"

	"github.com/variant64/server/entity"
)

type errorResponse struct {
	Error string `json:"error"`
}

// @Summary	Create a new player.
// @Accept	json
// @Produce	json
// @Router	/api/player [post]
// @Param	request	body		entity.RequestNewPlayer	true	"request body"
// @Success	200		{object}	entity.Player
// @Failure	400		{object}	errorResponse
func handlePostPlayer(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[entity.Player](w, req, &entity.RequestNewPlayer{})
}

// @Summary	Get player by id.
// @Produce	json
// @Router	/api/player/{player_id} [get]
// @Param	player_id	path		string	true	"player id"
// @Success	200		{object}	entity.Player
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handleGetPlayerByID(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[entity.Player](w, req, &entity.RequestGetPlayer{})
}

// @Summary	Create a new room.
// @Accept	json
// @Produce	json
// @Router	/api/room [post]
// @Param	request	body		entity.RequestNewRoom	true	"request body"
// @Success	200		{object}	entity.Room
// @Failure	400		{object}	errorResponse
func handlePostRoom(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[entity.Room](w, req, &entity.RequestNewRoom{})
}

// @Summary	Get all rooms.
// @Produce	json
// @Router	/api/rooms [get]
// @Success	200	{array}		entity.Room
// @Failure	404	{object}	errorResponse
// @Failure	500	{object}	errorResponse
func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	requestGetRooms := &entity.RequestGetRooms{}
	rooms := requestHandler.HandleGetRooms(requestGetRooms)

	writeEntityListResponse(w, rooms)
}

// @Summary	Get room by id.
// @Produce	json
// @Router	/api/room/{room_id} [get]
// @Param	room_id	path		string	true	"room id"
// @Success	200		{object}	entity.Room
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handleGetRoomByID(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[entity.Room](w, req, &entity.RequestGetRoom{})
}

// @Summary	Add player to a room.
// @Accept	json
// @Produce	json
// @Router	/api/room/{room_id}/join [post]
// @Param	room_id	path		string						true	"room id"
// @Param	request	body		entity.RequestJoinRoom	true	"request body"
// @Success	200		{object}	entity.Room
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[entity.Room](w, req, &entity.RequestJoinRoom{})
}

// @Summary	Remove player from a room.
// @Accept	json
// @Produce	json
// @Router	/api/room/{room_id}/leave [post]
// @Param	room_id	path		string							true	"room id"
// @Param	request	body		entity.RequestLeaveRoom	true	"request body"
// @Success	200		{object}	entity.Room
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostRoomLeave(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[entity.Room](w, req, &entity.RequestLeaveRoom{})
}

// @Summary	Start game in a room.
// @Accept	json
// @Produce	json
// @Router	/api/room/{room_id}/start [post]
// @Param	room_id	path		string					true	"room id"
// @Param	request	body		entity.RequestNewGame	true	"request body"
// @Success	200		{object}	entity.Room
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostRoomStart(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[entity.Room](w, req, &entity.RequestStartGame{})
}
