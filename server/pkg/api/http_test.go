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
	testEntities1 := Setup(
		WithPlayers(1),
	)

	testcases := []struct {
		description              string
		id                       string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid playerID.",
			testEntities1.player1.GetID().String(),
			[]string{
				fmt.Sprintf("\"display_name\":\"%s\"", testEntities1.player1.DisplayName),
				fmt.Sprintf("\"id\":\"%s\"", testEntities1.player1.GetID()),
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
	testEntities1 := Setup(
		WithRoom(),
	)

	testEntities2 := Setup(
		WithRoom(),
	)

	testcases := []struct {
		description              string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Multiple rooms.",
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", testEntities1.room1.Name),
				fmt.Sprintf("\"id\":\"%s\"", testEntities1.room1.GetID()),
				fmt.Sprintf("\"name\":\"%s\"", testEntities2.room1.Name),
				fmt.Sprintf("\"id\":\"%s\"", testEntities2.room1.GetID()),
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
	testEntities1 := Setup(
		WithRoom(),
	)

	testcases := []struct {
		description              string
		id                       string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID.",
			testEntities1.room1.ID.String(),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", testEntities1.room1.Name),
				fmt.Sprintf("\"id\":\"%s\"", testEntities1.room1.GetID()),
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
	testEntities1 := Setup(
		WithPlayers(1),
		WithRoom(),
	)

	testEntities2 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
	)

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID.",
			testEntities1.room1.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player1.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", testEntities1.room1.Name),
				fmt.Sprintf("\"id\":\"%s\"", testEntities1.room1.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", testEntities1.player1.GetID()),
				"\"game_id\":null",
			},
			200,
		},
		{
			"Valid room ID, but full.",
			testEntities2.room1.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", uuid.New()),
			[]string{"room has reached player_limit"},
			400,
		},
		{
			"Invalid body.",
			testEntities1.room1.ID.String(),
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			"1234",
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player1.GetID()),
			[]string{
				"failed to decode",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player1.GetID()),
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
	testEntities1 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
	)

	testEntities2 := Setup(
		WithPlayers(1),
		WithPlayersInRoom(1),
		WithRoom(),
	)

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID and playerID.",
			testEntities1.room1.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player1.GetID()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", testEntities1.room1.Name),
				fmt.Sprintf("\"id\":\"%s\"", testEntities1.room1.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", testEntities1.player2.GetID()),
				"\"game_id\":null",
			},
			200,
		},
		{
			"Valid room ID and invalid playerID.",
			testEntities2.room1.ID.String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", uuid.New()),
			[]string{
				fmt.Sprintf("\"name\":\"%s\"", testEntities2.room1.Name),
				fmt.Sprintf("\"id\":\"%s\"", testEntities2.room1.GetID()),
				fmt.Sprintf("\"players\":[\"%s\"]", testEntities2.player1.GetID()),
				"\"game_id\":null",
			},
			200,
		},
		{
			"Invalid body.",
			testEntities1.room1.ID.String(),
			"{",
			[]string{"failed to unmarshal request body"},
			400,
		},
		{
			"Invalid UUID.",
			"1234",
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player1.GetID()),
			[]string{
				"failed to decode",
			},
			400,
		},
		{
			"Invalid room ID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player1.GetID()),
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
	testEntities1 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
	)

	testEntities2 := Setup(
		WithPlayers(1),
		WithPlayersInRoom(1),
		WithRoom(),
	)

	testcases := []struct {
		description              string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid room ID and playerID.",
			fmt.Sprintf("{\"room_id\":\"%s\",\"player_time_ms\":1000000}", testEntities1.room1.GetID()),
			[]string{
				fmt.Sprintf("\"active_player\":\"%s\"", testEntities1.player1.GetID()),
				"\"state\":\"started\"",
			},
			200,
		},
		{
			"Valid room ID not enough players.",
			fmt.Sprintf("{\"room_id\":\"%s\",\"player_time_ms\":1000000}", testEntities2.room1.GetID()),
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
	testEntities1 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
		WithGame(),
	)

	testEntities2 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
		WithGame(),
		WithConcededPlayer(),
	)

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid gameID and playerID.",
			testEntities1.game1.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player1.GetID()),
			[]string{
				fmt.Sprintf("\"winning_players\":[\"%s\"]", testEntities1.player2.GetID()),
				fmt.Sprintf("\"losing_players\":[\"%s\"]", testEntities1.player1.GetID()),
			},
			200,
		},
		{
			"Valid gameID, but other player already conceded.",
			testEntities1.game1.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities2.player2.GetID()),
			[]string{"game is finished"},
			400,
		},
		{
			"Invalid gameID.",
			uuid.New().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player2.GetID()),
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
	testEntities1 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
		WithGame(),
		WithApprovedDrawPlayer(),
	)

	testEntities2 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
		WithGame(),
	)

	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid gameID and playerID, last to accept.",
			testEntities1.game1.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities1.player2.GetID()),
			[]string{
				"\"winning_players\":[]",
				"\"losing_players\":[]",
				fmt.Sprintf("\"drawn_players\":[\"%s\",\"%s\"]", testEntities1.player2.GetID(), testEntities1.player1.GetID()),
			},
			200,
		},
		{
			"Valid gameID and playerID, first to accept",
			testEntities2.game1.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\"}", testEntities2.player1.GetID()),
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
	testEntities1 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
		WithGame(),
	)

	requestApproveDrawPlayer1 := game.RequestApproveDraw{
		GameID:   testEntities1.game1.GetID(),
		PlayerID: testEntities1.player1.GetID(),
	}
	game1, err := requestApproveDrawPlayer1.PerformAction()
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

func TestGamePlayerMakeMove(t *testing.T) {
	testEntities1 := Setup(
		WithPlayers(2),
		WithPlayersInRoom(2),
		WithRoom(),
		WithGame(),
	)

	// Test making a valid move.
	testcases := []struct {
		description              string
		id                       string
		body                     string
		expectedResponseContains []string
		expectedStatusCode       int
	}{
		{
			"Valid gameID and move.",
			testEntities1.game1.GetID().String(),
			fmt.Sprintf("{\"player_id\":\"%s\",\"move\":{\"source\":{\"rank\":1,\"file\":1},\"destination\":{\"rank\":2,\"file\":1},\"move_type\":\"normal\"}}", testEntities1.player1.GetID()),
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
				fmt.Sprintf("/api/game/%s/move", tc.id),
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

type setupOption = func(s *setupBuilder)

type setupBuilder struct {
	player1          *player.Player
	player2          *player.Player
	addPlayer1InRoom bool
	addPlayer2InRoom bool
	room1            *room.Room
	game1            *game.Game
}

type testEntities struct {
	player1 *player.Player
	player2 *player.Player
	room1   *room.Room
	game1   *game.Game
}

func (s *setupBuilder) Build() testEntities {
	return testEntities{
		s.player1,
		s.player2,
		s.room1,
		s.game1,
	}
}

func Setup(options ...setupOption) testEntities {
	setupBuilder := &setupBuilder{}

	for _, option := range options {
		option(setupBuilder)
	}

	return setupBuilder.Build()
}

func WithPlayers(num int) setupOption {
	return func(s *setupBuilder) {
		if num <= 0 {
			panic("error in setup")
		}

		if num >= 1 {
			player1, err := (&player.RequestNewPlayer{DisplayName: "player1"}).PerformAction()
			if err != nil {
				panic("error in setup")
			}
			s.player1 = player1
		}

		if num >= 2 {
			player2, err := (&player.RequestNewPlayer{DisplayName: "player1"}).PerformAction()
			if err != nil {
				panic("error in setup")
			}
			s.player2 = player2
		}
	}
}

func WithPlayersInRoom(num int) setupOption {
	return func(s *setupBuilder) {
		if num <= 0 {
			panic("error in setup")
		}

		if num >= 1 {
			s.addPlayer1InRoom = true
		}

		if num >= 2 {
			s.addPlayer2InRoom = true
		}
	}
}

func WithRoom() setupOption {
	return func(s *setupBuilder) {
		room1, err := (&room.RequestNewRoom{Name: "test room"}).PerformAction()
		if err != nil {
			panic("error in setup")
		}

		if s.player1 != nil && s.addPlayer1InRoom {
			_, err = (&room.RequestJoinRoom{RoomID: room1.ID, PlayerID: s.player1.GetID()}).PerformAction()
			if err != nil {
				panic("error in setup")
			}
		}

		if s.player2 != nil && s.addPlayer2InRoom {
			_, err = (&room.RequestJoinRoom{RoomID: room1.ID, PlayerID: s.player2.GetID()}).PerformAction()
			if err != nil {
				panic("error in setup")
			}
		}

		s.room1 = room1
	}
}

func WithGame() setupOption {
	return func(s *setupBuilder) {
		game, err := (&room.RequestStartGame{RoomID: s.room1.GetID()}).PerformAction()
		if err != nil {
			panic("error in setup")
		}
		s.game1 = game
	}
}

func WithApprovedDrawPlayer() setupOption {
	return func(s *setupBuilder) {
		if s.game1 == nil {
			panic("error in setup")
		}

		_, err := (&game.RequestApproveDraw{
			GameID:   s.game1.GetID(),
			PlayerID: s.player1.GetID(),
		}).PerformAction()
		if err != nil {
			panic("error in setup")
		}
	}
}

func WithConcededPlayer() setupOption {
	return func(s *setupBuilder) {
		if s.game1 == nil {
			panic("error in setup")
		}

		_, err := (&game.RequestConcede{
			GameID:   s.game1.GetID(),
			PlayerID: s.player1.GetID(),
		}).PerformAction()
		if err != nil {
			panic("error in setup")
		}
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
