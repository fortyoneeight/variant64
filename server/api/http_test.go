package api

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/entity"
)

func TestPlayerPost(t *testing.T) {
	testcases := []struct {
		description              string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid user.",
			"{\"display_name\":\"user\"}",
			[]string{"\"display_name\":\"user\""},
			200,
		},
		{
			"Invalid user.",
			"{}",
			[]string{"display_name cannot be empty"},
			400,
		},
		{
			"Invalid body.",
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest("POST", "/api/player", strings.NewReader(tc.body))
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

func TestPlayerGetByID(t *testing.T) {
	playerName1 := "player1"
	player1 := &entity.Entity[entity.Player]{}
	requestNewPlayer1 := entity.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer1.Write(player1)
	player1.Store()

	testcases := []struct {
		description              string
		id                       string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid playerID.",
			player1.Data.GetID().String(),
			[]string{
				fmt.Sprintf("\"display_name\":\"%s\"", playerName1),
				fmt.Sprintf("\"id\":\"%s\"", player1.Data.GetID()),
			},
			200,
		},
		{
			"Invalid UUID.",
			"1234",
			[]string{
				"invalid id",
			},
			400,
		},
		{
			"Invalid playerID.",
			uuid.New().String(),
			[]string{
				"not found",
			},
			404,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest("GET", fmt.Sprintf("/api/player/%s", tc.id), nil)
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

func TestRoomPost(t *testing.T) {
	testcases := []struct {
		description              string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room.",
			"{\"room_name\":\"test\"}",
			[]string{
				"\"name\":\"test\"",
				"\"players\":[]",
			},
			200,
		},
		{
			"Invalid room.",
			"{}",
			[]string{
				"room_name cannot be empty",
			},
			400,
		},
		{
			"Invalid body.",
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest("POST", "/api/room", strings.NewReader(tc.body))
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

func TestRoomsGet(t *testing.T) {
	roomName1 := "room1"
	roomName2 := "room2"
	room1 := &entity.Entity[entity.Room]{}
	room2 := &entity.Entity[entity.Room]{}
	requestNewRoom1 := entity.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := entity.RequestNewRoom{Name: roomName2}
	requestNewRoom1.Write(room1)
	requestNewRoom2.Write(room2)
	room1.Store()
	room2.Store()

	testcases := []struct {
		description              string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Multiple rooms.",
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.Data.GetID()),
				fmt.Sprintf("\"name\":\"%s\"", roomName2),
				fmt.Sprintf("\"id\":\"%s\"", room2.Data.GetID()),
			},
			200,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest("GET", "/api/rooms", nil)
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

func TestRoomGetByID(t *testing.T) {
	roomName1 := "room1"
	room1 := &entity.Entity[entity.Room]{}
	requestNewRoom1 := entity.RequestNewRoom{Name: roomName1}
	requestNewRoom1.Write(room1)
	room1.Store()

	testcases := []struct {
		description              string
		id                       string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID.",
			room1.Data.ID.String(),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.Data.GetID()),
			},
			200,
		},
		{
			"Invalid UUID.",
			"1234",
			[]string{
				"invalid id",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			[]string{
				"not found",
			},
			404,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest("GET", fmt.Sprintf("/api/room/%s", tc.id), nil)
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

func TestRoomAddPlayer(t *testing.T) {
	playerName1 := "player1"
	player1 := &entity.Entity[entity.Player]{}
	requestNewPlayer1 := entity.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer1.Write(player1)
	player1.Store()

	roomName1 := "room1"
	room1 := &entity.Entity[entity.Room]{}
	requestNewRoom1 := entity.RequestNewRoom{Name: roomName1}
	requestNewRoom1.Write(room1)
	room1.Store()

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID.",
			room1.Data.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.Data.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.Data.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", player1.Data.GetID()),
			},
			200,
		},
		{
			"Invalid body.",
			room1.Data.ID.String(),
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			"1234",
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.Data.GetID()),
			[]string{
				"invalid id",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.Data.GetID()),
			[]string{
				"not found",
			},
			404,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest(
				"POST",
				fmt.Sprintf("/api/room/%s/join", tc.id),
				strings.NewReader(tc.body),
			)
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

func TestRoomRemovePlayer(t *testing.T) {
	playerName1 := "player1"
	playerName2 := "player2"
	player1 := &entity.Entity[entity.Player]{}
	player2 := &entity.Entity[entity.Player]{}
	requestNewPlayer1 := entity.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer2 := entity.RequestNewPlayer{DisplayName: playerName2}
	requestNewPlayer1.Write(player1)
	requestNewPlayer2.Write(player2)
	player1.Store()
	player2.Store()

	roomName1 := "room1"
	roomName2 := "room2"
	room1 := &entity.Entity[entity.Room]{}
	room2 := &entity.Entity[entity.Room]{}
	requestNewRoom1 := entity.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := entity.RequestNewRoom{Name: roomName2}
	requestNewRoom1.Write(room1)
	requestNewRoom2.Write(room2)

	requestAddPlayer1 := &entity.RequestRoomAddPlayer{
		PlayerID: player1.Data.GetID(),
	}

	requestAddPlayer1.RoomID = room1.Data.GetID()
	requestAddPlayer1.Write(room1)

	requestAddPlayer1.RoomID = room2.Data.GetID()
	requestAddPlayer1.Write(room2)

	requestAddPlayer2 := &entity.RequestRoomAddPlayer{
		PlayerID: player2.Data.GetID(),
	}

	requestAddPlayer1.RoomID = room1.Data.GetID()
	requestAddPlayer2.Write(room1)

	room1.Store()
	room2.Store()

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID and playerID.",
			room1.Data.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.Data.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.Data.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", player2.Data.GetID()),
			},
			200,
		},
		{
			"Valid room ID and invalid playerID.",
			room2.Data.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player2.Data.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName2),
				fmt.Sprintf("\"id\":\"%s\"", room2.Data.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", player1.Data.GetID()),
			},
			200,
		},
		{
			"Invalid body.",
			room1.Data.ID.String(),
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			"1234",
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.Data.GetID()),
			[]string{
				"invalid id",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.Data.GetID()),
			[]string{
				"not found",
			},
			404,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest(
				"POST",
				fmt.Sprintf("/api/room/%s/leave", tc.id),
				strings.NewReader(tc.body),
			)
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

func TestRoomStartGame(t *testing.T) {
	playerName1 := "player1"
	playerName2 := "player2"
	player1 := &entity.Entity[entity.Player]{}
	player2 := &entity.Entity[entity.Player]{}
	requestNewPlayer1 := entity.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer2 := entity.RequestNewPlayer{DisplayName: playerName2}
	requestNewPlayer1.Write(player1)
	requestNewPlayer2.Write(player2)
	player1.Store()
	player2.Store()

	roomName1 := "room1"
	roomName2 := "room2"
	room1 := &entity.Entity[entity.Room]{}
	room2 := &entity.Entity[entity.Room]{}
	requestNewRoom1 := entity.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := entity.RequestNewRoom{Name: roomName2}
	requestNewRoom1.Write(room1)
	requestNewRoom2.Write(room2)

	requestAddPlayer1 := &entity.RequestRoomAddPlayer{
		PlayerID: player1.Data.GetID(),
	}

	requestAddPlayer1.RoomID = room1.Data.GetID()
	requestAddPlayer1.Write(room1)

	requestAddPlayer1.RoomID = room2.Data.GetID()
	requestAddPlayer1.Write(room2)

	requestAddPlayer2 := &entity.RequestRoomAddPlayer{
		PlayerID: player2.Data.GetID(),
	}

	requestAddPlayer1.RoomID = room1.Data.GetID()
	requestAddPlayer2.Write(room1)

	room1.Store()
	room2.Store()

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID and playerID.",
			room1.Data.ID.String(),
			"{\"player_time_ms\":1000000}",
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"players\":[\"%s\",\"%s\"]", player1.Data.GetID(), player2.Data.GetID()),
			},
			200,
		},
		{
			"Valid room ID not enough players.",
			room2.Data.ID.String(),
			"{\"player_time_ms\":1000000}",
			[]string{
				"invalid number of players",
			},
			500,
		},
		{
			"Invalid body.",
			room1.Data.ID.String(),
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			"1234",
			"{\"player_time_ms\":1000000}",
			[]string{
				"invalid id",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			"{\"player_time_ms\":1000000}",
			[]string{
				"not found",
			},
			404,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest(
				"POST",
				fmt.Sprintf("/api/room/%s/start", tc.id),
				strings.NewReader(tc.body),
			)
			writer := executeRequest(router, request)

			assert.Equal(t, tc.expectedStatusCode, writer.statusCode)
			responseString := string(writer.response)
			for _, e := range tc.expectedResponseContains {
				assert.Contains(t, responseString, e)
			}
		})
	}
}

type mockWriter struct {
	response   []byte
	statusCode int
}

func (m *mockWriter) Header() http.Header {
	return nil
}

func (m *mockWriter) Write(bytes []byte) (int, error) {
	m.response = bytes
	return 0, nil
}

func (m *mockWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}

func executeRequest(router *mux.Router, request *http.Request) *mockWriter {
	mockWriter := &mockWriter{statusCode: 200, response: []byte{}}
	router.ServeHTTP(mockWriter, request)
	return mockWriter
}