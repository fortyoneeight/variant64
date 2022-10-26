package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	path    string
	handler func(w http.ResponseWriter, req *http.Request)
	methods []string
}

var httpRoutes = []route{
	{"/api/player", handlePostPlayer, []string{"POST"}},
	{"/api/player", handlePostPlayer, []string{"POST"}},
	{"/api/player/{id}", handleGetPlayerByID, []string{"GET"}},
	{"/api/room", handlePostRoom, []string{"POST"}},
	{"/api/rooms", handleGetRooms, []string{"GET"}},
	{"/api/room/{id}", handleGetRoomByID, []string{"GET"}},
	{"/api/room/{id}/join", handlePostRoomJoin, []string{"POST"}},
	{"/api/room/{id}/leave", handlePostRoomLeave, []string{"POST"}},
	{"/api/room/{id}/start", handlePostRoomStart, []string{"POST"}},
}

var websocketRoutes = []route{
	{"/ws", websocketHandler, []string{"GET"}},
}

func AttachRoutes(r *mux.Router) {
	for _, routeList := range [][]route{httpRoutes, websocketRoutes} {
		for _, route := range routeList {
			r.HandleFunc(route.path, route.handler).Methods(route.methods...)
		}
	}
}
