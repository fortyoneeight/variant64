package api

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/models/game"
	"github.com/variant64/server/pkg/models/player"
	"github.com/variant64/server/pkg/models/room"
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
			[]string{"missing display_name"},
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
	requestNewPlayer1 := player.RequestNewPlayer{DisplayName: playerName1}
	player1, err := requestNewPlayer1.PerformAction()
	assert.Nil(t, err)

	testcases := []struct {
		description              string
		id                       string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid playerID.",
			player1.GetID().String(),
			[]string{
				fmt.Sprintf("\"display_name\":\"%s\"", playerName1),
				fmt.Sprintf("\"id\":\"%s\"", player1.GetID()),
			},
			200,
		},
		{
			"Invalid UUID.",
			"1234",
			[]string{
				"failed to decode",
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
				"\"game_id\":null",
			},
			200,
		},
		{
			"Invalid room.",
			"{}",
			[]string{
				"room_name is required",
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
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := &room.RequestNewRoom{Name: roomName2}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)
	room2, err := requestNewRoom2.PerformAction()
	assert.Nil(t, err)

	testcases := []struct {
		description              string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Multiple rooms.",
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.GetID()),
				fmt.Sprintf("\"name\":\"%s\"", roomName2),
				fmt.Sprintf("\"id\":\"%s\"", room2.GetID()),
				"\"game_id\":null",
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
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)

	testcases := []struct {
		description              string
		id                       string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID.",
			room1.ID.String(),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.GetID()),
				"\"game_id\":null",
			},
			200,
		},
		{
			"Invalid UUID.",
			"1234",
			[]string{
				"failed to decode",
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
	requestNewPlayer1 := &player.RequestNewPlayer{DisplayName: playerName1}
	player1, err := requestNewPlayer1.PerformAction()
	assert.Nil(t, err)

	roomName1 := "room1"
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID.",
			room1.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", player1.GetID()),
				"\"game_id\":null",
			},
			200,
		},
		{
			"Invalid body.",
			room1.ID.String(),
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			"1234",
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
			[]string{
				"failed to decode",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
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
	requestNewPlayer1 := player.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer2 := player.RequestNewPlayer{DisplayName: playerName2}
	player1, err := requestNewPlayer1.PerformAction()
	assert.Nil(t, err)
	player2, err := requestNewPlayer2.PerformAction()
	assert.Nil(t, err)

	roomName1 := "room1"
	roomName2 := "room2"
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := &room.RequestNewRoom{Name: roomName2}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)
	room2, err := requestNewRoom2.PerformAction()
	assert.Nil(t, err)

	requestAddPlayer1 := &room.RequestJoinRoom{
		PlayerID: player1.GetID(),
	}

	requestAddPlayer1.RoomID = room1.GetID()
	room1, err = requestAddPlayer1.PerformAction()

	requestAddPlayer1.RoomID = room2.GetID()
	room2, err = requestAddPlayer1.PerformAction()

	requestAddPlayer2 := &room.RequestJoinRoom{
		PlayerID: player2.GetID(),
	}

	requestAddPlayer2.RoomID = room1.GetID()
	room1, err = requestAddPlayer2.PerformAction()

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID and playerID.",
			room1.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName1),
				fmt.Sprintf("\"id\":\"%s\"", room1.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", player2.GetID()),
				"\"game_id\":null",
			},
			200,
		},
		{
			"Valid room ID and invalid playerID.",
			room2.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player2.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", roomName2),
				fmt.Sprintf("\"id\":\"%s\"", room2.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", player1.GetID()),
				"\"game_id\":null",
			},
			200,
		},
		{
			"Invalid body.",
			room1.ID.String(),
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			"1234",
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
			[]string{
				"failed to decode",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
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
	requestNewPlayer1 := player.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer2 := player.RequestNewPlayer{DisplayName: playerName2}
	player1, err := requestNewPlayer1.PerformAction()
	assert.Nil(t, err)
	player2, err := requestNewPlayer2.PerformAction()
	assert.Nil(t, err)

	roomName1 := "room1"
	roomName2 := "room2"
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := &room.RequestNewRoom{Name: roomName2}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)
	room2, err := requestNewRoom2.PerformAction()
	assert.Nil(t, err)

	requestAddPlayer1 := &room.RequestJoinRoom{
		PlayerID: player1.GetID(),
	}

	requestAddPlayer1.RoomID = room1.GetID()
	room1, err = requestAddPlayer1.PerformAction()

	requestAddPlayer1.RoomID = room2.GetID()
	room2, err = requestAddPlayer1.PerformAction()

	requestAddPlayer2 := &room.RequestJoinRoom{
		PlayerID: player2.GetID(),
	}

	requestAddPlayer2.RoomID = room1.GetID()
	room1, err = requestAddPlayer2.PerformAction()

	testcases := []struct {
		description              string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID and playerID.",
			fmt.Sprintf("{\"room_id\":\"%s\",\"player_time_ms\":1000000}", room1.GetID()),
			[]string{
				fmt.Sprintf("\"active_player\":\"%s\"", player1.GetID()),
				"\"state\":\"started\"",
			},
			200,
		},
		{
			"Valid room ID not enough players.",
			fmt.Sprintf("{\"room_id\":\"%s\",\"player_time_ms\":1000000}", room2.GetID()),
			[]string{
				"invalid number of players",
			},
			400,
		},
		{
			"Invalid body.",
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			fmt.Sprintf("{\"room_id\":\"someid\",\"player_time_ms\":1000000}"),
			[]string{
				"failed to unmarshal request body",
			},
			400,
		},
		{
			"Invalid room ID.",
			fmt.Sprintf("{\"room_id\":\"%s\",\"player_time_ms\":1000000}", uuid.New()),
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
				fmt.Sprintf("/api/game"),
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

func TestGameConcede(t *testing.T) {
	playerName1 := "player1"
	playerName2 := "player2"
	requestNewPlayer1 := player.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer2 := player.RequestNewPlayer{DisplayName: playerName2}
	player1, err := requestNewPlayer1.PerformAction()
	assert.Nil(t, err)
	player2, err := requestNewPlayer2.PerformAction()
	assert.Nil(t, err)

	roomName1 := "room1"
	roomName2 := "room2"
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := &room.RequestNewRoom{Name: roomName2}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)
	room2, err := requestNewRoom2.PerformAction()
	assert.Nil(t, err)

	requestAddPlayer1 := &room.RequestJoinRoom{
		PlayerID: player1.GetID(),
	}

	requestAddPlayer1.RoomID = room1.GetID()
	room1, err = requestAddPlayer1.PerformAction()

	requestAddPlayer1.RoomID = room2.GetID()
	room2, err = requestAddPlayer1.PerformAction()

	requestAddPlayer2 := &room.RequestJoinRoom{
		PlayerID: player2.GetID(),
	}

	requestAddPlayer2.RoomID = room1.GetID()
	room1, err = requestAddPlayer2.PerformAction()

	requestAddPlayer2.RoomID = room2.GetID()
	room2, err = requestAddPlayer2.PerformAction()

	requestStartGame1 := &room.RequestStartGame{
		RoomID:          room1.GetID(),
		PlayerTimeMilis: 1_000_000,
	}
	game1, err := requestStartGame1.PerformAction()

	requestStartGame2 := &room.RequestStartGame{
		RoomID:          room2.GetID(),
		PlayerTimeMilis: 1_000_000,
	}
	game2, err := requestStartGame2.PerformAction()

	requestConcede := &game.RequestConcede{
		GameID:   game2.GetID(),
		PlayerID: player1.ID,
	}
	requestConcede.PerformAction()

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid gameID and playerID.",
			game1.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
			[]string{
				fmt.Sprintf("\"winning_players\":[\"%s\"]", player2.GetID()),
				fmt.Sprintf("\"losing_players\":[\"%s\"]", player1.GetID()),
			},
			200,
		},
		{
			"Valid gameID, but other player already conceded.",
			game2.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player2.GetID()),
			[]string{"game is finished"},
			400,
		},
		{
			"Invalid gameID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player2.GetID()),
			[]string{"not found"},
			404,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest(
				"POST",
				fmt.Sprintf("/api/game/%s/concede", tc.id),
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

func TestGameDrawApprove(t *testing.T) {
	playerName1 := "player1"
	playerName2 := "player2"
	requestNewPlayer1 := player.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer2 := player.RequestNewPlayer{DisplayName: playerName2}
	player1, err := requestNewPlayer1.PerformAction()
	assert.Nil(t, err)
	player2, err := requestNewPlayer2.PerformAction()
	assert.Nil(t, err)

	roomName1 := "room1"
	roomName2 := "room2"
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	requestNewRoom2 := &room.RequestNewRoom{Name: roomName2}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)
	room2, err := requestNewRoom2.PerformAction()
	assert.Nil(t, err)

	requestAddPlayer1 := &room.RequestJoinRoom{
		PlayerID: player1.GetID(),
	}

	requestAddPlayer1.RoomID = room1.GetID()
	room1, err = requestAddPlayer1.PerformAction()

	requestAddPlayer1.RoomID = room2.GetID()
	room2, err = requestAddPlayer1.PerformAction()

	requestAddPlayer2 := &room.RequestJoinRoom{
		PlayerID: player2.GetID(),
	}

	requestAddPlayer2.RoomID = room1.GetID()
	room1, err = requestAddPlayer2.PerformAction()

	requestAddPlayer2.RoomID = room2.GetID()
	room2, err = requestAddPlayer2.PerformAction()

	requestStartGame1 := &room.RequestStartGame{
		RoomID:          room1.GetID(),
		PlayerTimeMilis: 1_000_000,
	}
	game1, err := requestStartGame1.PerformAction()
	assert.Nil(t, err)

	requestApproveDrawPlayer1 := game.RequestApproveDraw{
		GameID:   game1.GetID(),
		PlayerID: player1.GetID(),
	}
	game1, err = requestApproveDrawPlayer1.PerformAction()
	assert.Nil(t, err)

	requestStartGame2 := &room.RequestStartGame{
		RoomID:          room2.GetID(),
		PlayerTimeMilis: 1_000_000,
	}
	game2, err := requestStartGame2.PerformAction()
	assert.Nil(t, err)

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid gameID and playerID, last to accept.",
			game1.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player2.GetID()),
			[]string{
				"\"winning_players\":[]",
				"\"losing_players\":[]",
				fmt.Sprintf("\"drawn_players\":[\"%s\",\"%s\"]", player2.GetID(), player1.GetID()),
			},
			200,
		},
		{
			"Valid gameID and playerID, first to accept",
			game2.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", player1.GetID()),
			[]string{
				"\"winning_players\":[]",
				"\"losing_players\":[]",
				"\"drawn_players\":[]",
			},
			200,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest(
				"POST",
				fmt.Sprintf("/api/game/%s/draw/approve", tc.id),
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

func TestGameDrawReject(t *testing.T) {
	playerName1 := "player1"
	playerName2 := "player2"
	requestNewPlayer1 := player.RequestNewPlayer{DisplayName: playerName1}
	requestNewPlayer2 := player.RequestNewPlayer{DisplayName: playerName2}
	player1, err := requestNewPlayer1.PerformAction()
	assert.Nil(t, err)
	player2, err := requestNewPlayer2.PerformAction()
	assert.Nil(t, err)

	roomName1 := "room1"
	requestNewRoom1 := &room.RequestNewRoom{Name: roomName1}
	room1, err := requestNewRoom1.PerformAction()
	assert.Nil(t, err)

	requestAddPlayer1 := &room.RequestJoinRoom{
		PlayerID: player1.GetID(),
	}

	requestAddPlayer1.RoomID = room1.GetID()
	room1, err = requestAddPlayer1.PerformAction()

	requestAddPlayer2 := &room.RequestJoinRoom{
		PlayerID: player2.GetID(),
	}

	requestAddPlayer2.RoomID = room1.GetID()
	room1, err = requestAddPlayer2.PerformAction()

	requestStartGame1 := &room.RequestStartGame{
		RoomID:          room1.GetID(),
		PlayerTimeMilis: 1_000_000,
	}
	game1, err := requestStartGame1.PerformAction()
	assert.Nil(t, err)

	requestApproveDrawPlayer1 := game.RequestApproveDraw{
		GameID:   game1.GetID(),
		PlayerID: player1.GetID(),
	}
	game1, err = requestApproveDrawPlayer1.PerformAction()
	assert.Nil(t, err)

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid gameID.",
			game1.GetID().String(),
			"{}",
			[]string{
				"\"winning_players\":[]",
				"\"losing_players\":[]",
				"\"drawn_players\":[]",
			},
			200,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			router := &mux.Router{}
			AttachRoutes(router)

			request, _ := http.NewRequest(
				"POST",
				fmt.Sprintf("/api/game/%s/draw/reject", tc.id),
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
