package api

import (
	"net/http"

	"github.com/variant64/server/game"
	"github.com/variant64/server/player"
	"github.com/variant64/server/room"
)

type errorResponse struct {
	Error string `json:"error"`
}

// @Summary	Create a new player.
// @Accept	json
// @Produce	json
// @Router	/api/player [post]
// @Param	request	body		player.RequestNewPlayer	true	"request body"
// @Success	200		{object}	player.Player
// @Failure	400		{object}	errorResponse
func handlePostPlayer(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*player.Player](w, req, &player.RequestNewPlayer{})
}

// @Summary	Get player by id.
// @Produce	json
// @Router	/api/player/{player_id} [get]
// @Param	player_id	path		string	true	"player id"
// @Success	200		{object}	player.Player
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handleGetPlayerByID(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*player.Player](w, req, &player.RequestGetPlayer{})
}

// @Summary	Create a new room.
// @Accept	json
// @Produce	json
// @Router	/api/room [post]
// @Param	request	body		room.RequestNewRoom	true	"request body"
// @Success	200		{object}	room.Room
// @Failure	400		{object}	errorResponse
func handlePostRoom(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*room.Room](w, req, &room.RequestNewRoom{})
}

// @Summary	Get all rooms.
// @Produce	json
// @Router	/api/rooms [get]
// @Success	200	{array}		room.Room
// @Failure	404	{object}	errorResponse
// @Failure	500	{object}	errorResponse
func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[[]*room.Room](w, req, &room.RequestGetRooms{})
}

// @Summary	Get room by id.
// @Produce	json
// @Router	/api/room/{room_id} [get]
// @Param	room_id	path		string	true	"room id"
// @Success	200		{object}	room.Room
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handleGetRoomByID(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*room.Room](w, req, &room.RequestGetRoom{})
}

// @Summary	Add player to a room.
// @Accept	json
// @Produce	json
// @Router	/api/room/{room_id}/join [post]
// @Param	room_id	path		string						true	"room id"
// @Param	request	body		room.RequestJoinRoom	    true	"request body"
// @Success	200		{object}	room.Room
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*room.Room](w, req, &room.RequestJoinRoom{})
}

// @Summary	Remove player from a room.
// @Accept	json
// @Produce	json
// @Router	/api/room/{room_id}/leave [post]
// @Param	room_id	path		string							true	"room id"
// @Param	request	body		room.RequestLeaveRoom	true	"request body"
// @Success	200		{object}	room.Room
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostRoomLeave(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*room.Room](w, req, &room.RequestLeaveRoom{})
}

// @Summary	Start a game.
// @Accept	json
// @Produce	json
// @Router	/api/game [post]
// @Param	request	body		room.RequestStartGame	true	"request body"
// @Success	200		{object}	game.Game
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostGame(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*game.Game](w, req, &room.RequestStartGame{})
}

// @Summary	Player concedes a game.
// @Accept	json
// @Produce	json
// @Router	/api/game/{game_id}/concede [post]
// @Param	game_id	path		string					true	"room id"
// @Param	request	body		game.RequestConcede	    true	"request body"
// @Success	200		{object}	game.Game
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostGamePlayerConcede(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*game.Game](w, req, &game.RequestConcede{})
}

// @Summary	Player approves a game to draw.
// @Accept	json
// @Produce	json
// @Router	/api/game/{game_id}/draw/approve [post]
// @Param	game_id	path		string					true	"room id"
// @Param	request	body		game.RequestApproveDraw true	"request body"
// @Success	200		{object}	game.Game
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostGamePlayerApproveDraw(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*game.Game](w, req, &game.RequestApproveDraw{})
}

// @Summary	Player approves a game to draw.
// @Accept	json
// @Produce	json
// @Router	/api/game/{game_id}/draw/reject [post]
// @Param	game_id	path		string					true	"room id"
// @Param	request	body		game.RequestRejectDraw	true	"request body"
// @Success	200		{object}	game.Game
// @Failure	400		{object}	errorResponse
// @Failure	404		{object}	errorResponse
// @Failure	500		{object}	errorResponse
func handlePostGamePlayerRejectDraw(w http.ResponseWriter, req *http.Request) {
	handleActionRoute[*game.Game](w, req, &game.RequestRejectDraw{})
}
