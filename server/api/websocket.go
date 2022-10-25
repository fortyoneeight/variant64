package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/variant64/server/api/command"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// commandHandler represents a handler of incoming client commands.
type commandHandler interface {
	HandleCommand(command command.Command, message []byte) error
}

// websocketHandler upgrades and handles an incoming websocket connection request.
func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("error during connection upgrade: ", err)
		return
	}
	defer conn.Close()

	handler := command.NewWSHandler(conn)
	readAndHandleMessages(conn, handler)
}

// readAndHandleMessages continuously reads client messages and handles them.
func readAndHandleMessages(conn *websocket.Conn, handler commandHandler) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("error during message reading: ", err)
			break
		}

		command := &command.Command{}
		err = json.Unmarshal(message, command)
		if err == nil {
			handler.HandleCommand(*command, message)
		}
	}
}
