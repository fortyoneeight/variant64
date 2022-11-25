package entity

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type RequestHandler struct{}

// HandleNewPlayer handles a RequestNewPlayer.
func (h *RequestHandler) HandleNewPlayer(r *RequestNewPlayer) (*Entity[Player], error) {
	player := &Entity[Player]{}
	err := r.Write(player)
	if err != nil {
		return nil, err
	}

	player.Store()

	return player, nil
}

// HandleGetPlayer handles a RequestGetPlayer.
func (h *RequestHandler) HandleGetPlayer(r *RequestGetPlayer) (*Entity[Player], error) {
	player := &Entity[Player]{}
	err := r.Read(player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

// HandleNewRoom handles a RequestNewRoom.
func (h *RequestHandler) HandleNewRoom(r *RequestNewRoom) (*Entity[Room], error) {
	room := &Entity[Room]{}
	err := r.Write(room)
	if err != nil {
		return nil, err
	}

	room.Store()

	return room, nil
}

// HandleGetRooms handles a RequestGetRooms.
func (h *RequestHandler) HandleGetRooms(r *RequestGetRooms) *EntityList[Room] {
	rooms := &EntityList[Room]{}
	r.Read(rooms)
	return rooms
}

// HandleGetRoom handles a RequestGetRoom.
func (h *RequestHandler) HandleGetRoom(r *RequestGetRoom) (*Entity[Room], error) {
	room := &Entity[Room]{}
	err := r.Read(room)
	if err != nil {
		return nil, err
	}
	return room, nil
}

// RequestRoomAddPlayer is used to add a Player to a Room.
type RequestRoomAddPlayer struct {
	RoomID   uuid.UUID `json:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

// HandleRoomAddPlayer handles a RequestRoomAddPlayer.
func (h *RequestHandler) HandleRoomAddPlayer(r *RequestRoomAddPlayer) (*Entity[Room], error) {
	requestGetRoom := &RequestGetRoom{ID: r.RoomID}
	room, err := h.HandleGetRoom(requestGetRoom)
	if err != nil || room == nil {
		return nil, err
	}

	for _, p := range room.Data.Players {
		if p == r.PlayerID {
			return nil, errors.New("player cannot be duplicate")
		}
	}
	room.Data.Players = append(room.Data.Players, r.PlayerID)

	room.Store()

	return room, nil
}

// RequestRoomRemovePlayer is used to remove a Player from a Room.
type RequestRoomRemovePlayer struct {
	RoomID   uuid.UUID `json:"room_id"`
	PlayerID uuid.UUID `json:"player_id"`
}

// HandleRoomRemovePlayer handles a RequestRoomAddPlayer.
func (h *RequestHandler) HandleRoomRemovePlayer(r *RequestRoomRemovePlayer) (*Entity[Room], error) {
	requestGetRoom := &RequestGetRoom{ID: r.RoomID}
	room, err := h.HandleGetRoom(requestGetRoom)
	if err != nil {
		return nil, err
	}

	for i, p := range room.Data.Players {
		if p == r.PlayerID {
			room.Data.Players = append(room.Data.Players[:i], room.Data.Players[i+1:]...)
		}
	}

	room.Store()

	return room, nil
}

// HandleNewGame handles a RequestNewGame.
func (h *RequestHandler) HandleNewGame(r *RequestNewGame) (*Entity[*Game], error) {
	game := &Entity[*Game]{}
	err := r.Write(game)
	if err != nil {
		return nil, err
	}
	game.Store()
	return game, nil
}

// RequestRoomStartGame is used to start a new Game in a Room.
type RequestRoomStartGame struct {
	RoomID          uuid.UUID `json:"room_id"`
	PlayerTimeMilis int64     `json:"player_time_ms"`
}

// HandleRequestGameStart handles a RequestGameStart.
func (h *RequestHandler) HandleRoomStartGame(r *RequestRoomStartGame) (*Entity[Room], error) {
	requestGetRoom := &RequestGetRoom{
		ID: r.RoomID,
	}
	room, err := h.HandleGetRoom(requestGetRoom)
	if err != nil {
		return nil, err
	}

	requestNewGame := &RequestNewGame{
		PlayerOrder:     room.Data.Players,
		PlayerTimeMilis: 1_000 * 60 * 10,
	}
	game, err := h.HandleNewGame(requestNewGame)
	if err != nil {
		return nil, err
	}

	gameID := game.Data.GetID()
	room.Data.GameID = &gameID

	game.Data.start()

	room.Store()
	game.Store()

	return room, nil
}
