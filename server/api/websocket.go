package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func AttachWebsocketRoutes(r *mux.Router) {
	r.HandleFunc("/ws", websocketHandler).Methods("GET")
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("error during connection upgrade: ", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("error during message reading: ", err)
			break
		}

		responses := mockResponseHandler(message)
		for _, resp := range responses {
			err = conn.WriteMessage(messageType, []byte(resp))
			if err != nil {
				log.Println("Error during message writing:", err)
				break
			}
		}
	}
}

func mockResponseHandler(message []byte) []string {
	messageJSON := make(map[string]string)
	_ = json.Unmarshal(message, &messageJSON)

	if val, ok := messageJSON["command"]; ok {
		switch val {
		case "subscribe":
			return []string{
				"{\"channel\":\"room\",\"room_name\":\"roomname\",\"type\":\"subscribe\"}",
				"{\"channel\":\"room\",\"room_name\":\"roomname\",\"type\":\"snapshot\",\"data\":{\"board\": {\"size\":{\"length\":10,\"width\":10},\"cells\":[{\"x\":0,\"y\":0,\"cellItem\":{\"type\":\"piece\",\"data\":{\"name\":\"pawn\",player:{\"id\":\"uuid\",\"display_name\":\"player1\"},\"moves\":[\"a1\", \"b2\", \"c3\"]}}}]},\"active_player\":\"player_1\",\"clocks\":{\"player1\":100,\"player2\":100}}",
			}
		case "unsubscribe":
			return []string{"{\"channel\":\"room\",\"room_name\":\"roomname\",\"type\":\"unsubscribe\"}"}
		default:
			return []string{"{\"error\":\"unknown command\"}"}
		}
	}
	return []string{"{\"error\":\"unknown command\"}"}
}
