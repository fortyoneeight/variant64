package api

import (
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
	{"/", "Return status.", handleStatus, []string{"GET"}},
	{"/api/player", "Create a Player.", handlePostPlayer, []string{"POST"}},
	{"/api/player/{id}", "Get a Player by ID.", handleGetPlayerByID, []string{"GET"}},
	{"/api/room", "Create a Room.", handlePostRoom, []string{"POST"}},
	{"/api/rooms", "Get all Rooms.", handleGetRooms, []string{"GET"}},
	{"/api/room/{id}", "Get a Room by ID.", handleGetRoomByID, []string{"GET"}},
	{"/api/room/{id}/join", "Add a Player to a Room.", handlePostRoomJoin, []string{"POST"}},
	{"/api/room/{id}/leave", "Remove a Player from a Room.", handlePostRoomLeave, []string{"POST"}},
	{"/api/room/{id}/start", "Start the Game in a Room.", handlePostRoomStart, []string{"POST"}},
}

var websocketRoutes = []route{
	{"/ws", "Open a bi-directional websocket connection.", websocketHandler, []string{"GET"}},
}

// AttachRoutes adds all server routes to the provided mux.Router.
func AttachRoutes(r *mux.Router) {
	for _, routeList := range [][]route{httpRoutes, websocketRoutes} {
		for _, route := range routeList {
			r.HandleFunc(route.path, route.handler).Methods(route.methods...)
		}
	}
}
