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

func logRequest(w http.ResponseWriter, req *http.Request) {
	logger.Info(fmt.Sprintf("[HANDLING_ROUTE] %s %s %s\n", req.RemoteAddr, req.Method, req.URL))
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
