package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	path        string
	description string
	handler     func(w http.ResponseWriter, req *http.Request)
	methods     []string
}

var httpRoutes = []route{
	{"/api/player", "Create a Player.", handlePostPlayer, []string{"POST"}},
	{"/api/player/{player_id}", "Get a Player by ID.", handleGetPlayerByID, []string{"GET"}},
	{"/api/room", "Create a Room.", handlePostRoom, []string{"POST"}},
	{"/api/rooms", "Get all Rooms.", handleGetRooms, []string{"GET"}},
	{"/api/room/{room_id}", "Get a Room by ID.", handleGetRoomByID, []string{"GET"}},
	{"/api/room/{room_id}/join", "Add a Player to a Room.", handlePostRoomJoin, []string{"POST"}},
	{"/api/room/{room_id}/leave", "Remove a Player from a Room.", handlePostRoomLeave, []string{"POST"}},
	{"/api/game", "Start the Game.", handlePostGame, []string{"POST"}},
	{"/api/game/{game_id}/concede", "Player concedes a Game.", handlePostGamePlayerConcede, []string{"POST"}},
	{"/api/game/{game_id}/draw/approve", "Player approves a drawn Game.", handlePostGamePlayerApproveDraw, []string{"POST"}},
	{"/api/game/{game_id}/draw/reject", "Player rejects a drawn Game.", handlePostGamePlayerRejectDraw, []string{"POST"}},
	{"/api/game/{game_id}/move", "Player makes a move in the game.", handlePostGamePlayerMakeMove, []string{"POST"}},
}

var websocketRoutes = []route{
	{"/ws", "Open a bi-directional websocket connection.", websocketHandler, []string{"GET"}},
}

func logRequest(w http.ResponseWriter, req *http.Request) {
	logger.Info(fmt.Sprintf("[HANDLING_ROUTE] %s %s %s", req.RemoteAddr, req.Method, req.URL))
}

func createHandler(handler http.HandlerFunc) http.HandlerFunc {
	handlers := []http.HandlerFunc{
		logRequest,
		handler,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for _, call := range handlers {
			call(w, r)
		}
	}
}

// AttachRoutes adds all server routes to the provided mux.Router.
func AttachRoutes(r *mux.Router) {
	for _, routeList := range [][]route{httpRoutes, websocketRoutes} {
		for _, route := range routeList {
			r.HandleFunc(route.path, createHandler(route.handler)).Methods(route.methods...)
		}
	}
}
