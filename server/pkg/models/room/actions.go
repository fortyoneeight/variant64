package room

import (
	"encoding/json"
	"sync"

	"github.com/google/uuid"
	"github.com/variant64/server/pkg/models"
	"github.com/variant64/server/pkg/models/game"
)

var roomUpdateBus = models.NewUpdateBus[RoomUpdate]()

// RequestNewRoom is used to create a new Room.
type RequestNewRoom struct {
	Name string `json:"room_name" mapstructure:"room_name"`
}

// PerformAction creates a new Room.
func (r *RequestNewRoom) PerformAction() (*Room, error) {
	if r.Name == "" {
		return nil, errMissingName
	}

	room := &Room{
		ID:      uuid.New(),
		Name:    r.Name,
		Players: make([]uuid.UUID, 0),
		mux:     &sync.RWMutex{},
	}

	handler, err := models.NewUpdatePub(room.ID, roomUpdateBus)
	if err != nil {
		return nil, models.ErrFailedUpdatePub("Room")
	}
	room.updateHandler = handler

	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	roomStore.Store(room)

	return room, nil
}

// RequestGetRoom is used to get a Room by its ID.
type RequestGetRoom struct {
	RoomID uuid.UUID `json:"room_id" mapstructure:"room_id"`
}

// PerformAction loads a Room.
func (r *RequestGetRoom) PerformAction() (*Room, error) {
	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	room := roomStore.GetByID(r.RoomID)
	if room == nil {
		return nil, errRoomNotFound
	}

	return room, nil
}

// RequestGetRooms is used to get all Rooms.
type RequestGetRooms struct{}

// PerformAction gets all rooms.
func (r *RequestGetRooms) PerformAction() ([]*Room, error) {
	roomStore := getRoomStore()
	roomStore.Lock()
	defer roomStore.Unlock()

	rooms := roomStore.GetAll()
	if rooms == nil {
		return nil, errRoomNotFound
	}

	return rooms, nil
}

// RequestJoinRoom is used to add a Player to a Room.
type RequestJoinRoom struct {
	RoomID   uuid.UUID `json:"room_id" mapstructure:"room_id" swaggerignore:"true"`
	PlayerID uuid.UUID `json:"player_id"`
}

func (r *RequestJoinRoom) PerformAction() (*Room, error) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	room.mux.Lock()
	defer room.mux.Unlock()

	for _, p := range room.Players {
		if p == r.PlayerID {
			return nil, errDuplicatePlayer(r.PlayerID.String())
		}
	}
	room.Players = append(room.Players, r.PlayerID)

	room.updateHandler.Publish(
		models.UpdateMessage[RoomUpdate]{
			Channel: MessageChannel,
			Type:    models.UpdateType_DELTA,
			Data: RoomUpdate{
				ID:      &r.RoomID,
				Players: &room.Players,
			},
		},
	)

	return room, nil
}

// RequestLeaveRoom is used to remove a Player from a Room.
type RequestLeaveRoom struct {
	RoomID   uuid.UUID `json:"room_id" mapstructure:"room_id" swaggerignore:"true"`
	PlayerID uuid.UUID `json:"player_id"`
}

// RequestLeaveRoom handles a RequestRoomAddPlayer.
func (r *RequestLeaveRoom) PerformAction() (*Room, error) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	room.mux.Lock()
	defer room.mux.Unlock()

	for i, p := range room.Players {
		if p == r.PlayerID {
			room.Players = append(room.Players[:i], room.Players[i+1:]...)
		}
	}

	room.updateHandler.Publish(
		models.UpdateMessage[RoomUpdate]{
			Channel: MessageChannel,
			Type:    models.UpdateType_DELTA,
			Data: RoomUpdate{
				ID:      &r.RoomID,
				Players: &room.Players,
			},
		},
	)

	return room, nil
}

// RequestStartGame is used to start a new Game in a Room.
type RequestStartGame struct {
	RoomID          uuid.UUID `json:"room_id" mapstructure:"room_id"`
	PlayerTimeMilis int64     `json:"player_time_ms"`
}

// PerformAction starts a game.Game in a Room.
func (r *RequestStartGame) PerformAction() (*game.Game, error) {
	room, err := (&RequestGetRoom{RoomID: r.RoomID}).PerformAction()
	if err != nil || room == nil {
		return nil, err
	}

	room.mux.Lock()
	defer room.mux.Unlock()

	gameEntity, err := (&game.RequestNewGame{
		PlayerOrder:     room.Players,
		PlayerTimeMilis: r.PlayerTimeMilis,
	}).PerformAction()
	if err != nil || gameEntity == nil {
		return nil, err
	}

	gameEntity, err = (&game.RequestStartGame{
		GameID: gameEntity.GetID(),
	}).PerformAction()
	if err != nil || gameEntity == nil {
		return nil, err
	}

	gameID := gameEntity.GetID()
	room.GameID = &gameID

	room.updateHandler.Publish(
		models.UpdateMessage[RoomUpdate]{
			Channel: MessageChannel,
			Type:    models.UpdateType_DELTA,
			Data: RoomUpdate{
				ID:      &r.RoomID,
				Players: &room.Players,
			},
		},
	)

	return gameEntity, nil
}

const (
	MessageChannel = "room"

	RoomSubscribe   string = "subscribe"
	RoomUnsubscribe string = "unsubscribe"
)

// CommandRoomSubscribe represents a room subscribe command.
type CommandRoomSubscribe struct {
	RoomID      uuid.UUID `json:"room_id"`
	EventWriter models.EventWriter
}

func (c *CommandRoomSubscribe) PerformAction() error {
	room, err := (&RequestGetRoom{RoomID: c.RoomID}).PerformAction()
	if err != nil {
		return err
	}

	// Subscribe to updates.
	models.Subscribe(roomUpdateBus, c.RoomID, MessageChannel, c.EventWriter)

	// Send Room snapshot upon subscription.
	snapshot := models.UpdateMessage[RoomUpdate]{
		Channel: MessageChannel,
		Type:    models.UpdateType_SNAPSHOT,
		Data:    room.getSnapshot(),
	}
	snapshotBytes, err := json.Marshal(snapshot)
	c.EventWriter.WriteMessage(1, snapshotBytes)

	return nil
}

// CommandRoomUnsubscribe represents an room unsubscribe command.
type CommandRoomUnsubscribe struct {
	models.Command
	RoomID uuid.UUID `json:"room_id"`
}

func (c *CommandRoomUnsubscribe) PerformAction() error {
	return nil
}

// HandleCommand handles all incoming room websocket messages.
func HandleCommand(writer models.EventWriter, command, body string) error {
	switch {
	case command == RoomSubscribe:
		return models.HandleCommand(models.MarshallCommand(body, &CommandRoomSubscribe{EventWriter: writer}))
	case command == RoomUnsubscribe:
		c := CommandRoomUnsubscribe{}
		return c.PerformAction()
	default:
		return models.ErrInvalidCommand
	}
}
