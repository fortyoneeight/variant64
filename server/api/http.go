package api

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/variant64/server/entity"
	"github.com/variant64/server/store"
)

func handlePostPlayer(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	player := &entity.Entity[entity.Player]{}

	requestNewPlayer := &entity.RequestNewPlayer{}
	err = json.Unmarshal(body, requestNewPlayer)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}
	err = requestNewPlayer.Write(player)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	player.Store()

	writeEntityResponse(w, player)
}

func handleGetPlayerByID(w http.ResponseWriter, req *http.Request) {
	id := parseIDFromVars(req)
	if !id.Valid {
		writeBadRequestResponse(w, errors.New("invalid id"))
		return
	}

	player := &entity.Entity[entity.Player]{}
	requestGetPlayer := &entity.RequestGetPlayer{ID: id.UUID}
	err := requestGetPlayer.Read(player)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	err = player.Load()
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, player)
}

func handlePostRoom(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	room := &entity.Entity[entity.Room]{}

	requestNewRoom := &entity.RequestNewRoom{}
	err = json.Unmarshal(body, requestNewRoom)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}
	err = requestNewRoom.Write(room)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	room.Store()

	writeEntityResponse(w, room)
}

func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	rooms := &entity.EntityList[entity.Room]{}

	requestGetRooms := &entity.RequestGetRooms{}
	err := requestGetRooms.Read(rooms)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	rooms.Load()

	writeEntityListResponse(w, rooms)
}

func handleGetRoomByID(w http.ResponseWriter, req *http.Request) {
	id := parseIDFromVars(req)
	if !id.Valid {
		writeBadRequestResponse(w, errors.New("invalid id"))
		return
	}

	room := &entity.Entity[entity.Room]{}

	requestGetRoom := &entity.RequestGetRoom{ID: id.UUID}
	err := requestGetRoom.Read(room)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	err = room.Load()
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, room)
}

func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	id := parseIDFromVars(req)
	if !id.Valid {
		writeBadRequestResponse(w, errors.New("invalid id"))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	room := &entity.Entity[entity.Room]{}

	requestGetRoom := &entity.RequestGetRoom{ID: id.UUID}
	err = requestGetRoom.Read(room)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	err = room.Load()
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	requestAddPlayer := &entity.RequestRoomAddPlayer{RoomID: id.UUID}
	err = json.Unmarshal(body, requestAddPlayer)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}
	err = requestAddPlayer.Write(room)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	room.Store()

	writeEntityResponse(w, room)
}

func handlePostRoomLeave(w http.ResponseWriter, req *http.Request) {
	id := parseIDFromVars(req)
	if !id.Valid {
		writeBadRequestResponse(w, errors.New("invalid id"))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	room := &entity.Entity[entity.Room]{}

	requestGetRoom := &entity.RequestGetRoom{ID: id.UUID}
	err = requestGetRoom.Read(room)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	err = room.Load()
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	requestRemovePlayer := &entity.RequestRoomRemovePlayer{RoomID: id.UUID}
	err = json.Unmarshal(body, requestRemovePlayer)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}
	err = requestRemovePlayer.Write(room)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	room.Store()

	writeEntityResponse(w, room)
}

func handlePostRoomStart(w http.ResponseWriter, req *http.Request) {
	id := parseIDFromVars(req)
	if !id.Valid {
		writeBadRequestResponse(w, errors.New("invalid id"))
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	room := &entity.Entity[entity.Room]{}

	requestGetRoom := &entity.RequestGetRoom{ID: id.UUID}
	err = requestGetRoom.Read(room)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	err = room.Load()
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	game := &entity.Entity[*entity.Game]{}

	requestNewGame := &entity.RequestNewGame{PlayerOrder: room.Data.Players}
	err = json.Unmarshal(body, requestNewGame)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}
	err = requestNewGame.Write(game)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	requestGameStart := &entity.RequestGameStart{}
	err = requestGameStart.Write(game)
	if err != nil {
		writeServerErrorResponse(w, err)
		return
	}

	game.Store()

	writeEntityResponse(w, room)
}

func parseIDFromVars(req *http.Request) uuid.NullUUID {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		return uuid.NullUUID{Valid: false}
	}
	return uuid.NullUUID{UUID: id, Valid: true}
}

func writeBadRequestResponse(w http.ResponseWriter, err error) {
	logger.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
}

func writeNotFoundResponse(w http.ResponseWriter, err error) {
	logger.Error(err)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
}

func writeServerErrorResponse(w http.ResponseWriter, err error) {
	logger.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
}

func writeEntityResponse[T store.Indexable](w http.ResponseWriter, e *entity.Entity[T]) {
	serialized, err := json.Marshal(e.Data)
	if err != nil {
		writeServerErrorResponse(w, err)
	}
	w.Write(serialized)
}

func writeEntityListResponse[T store.Indexable](w http.ResponseWriter, e *entity.EntityList[T]) {
	serialized, _ := json.Marshal(e.Data)
	w.Write(serialized)
}
