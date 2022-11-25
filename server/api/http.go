package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/variant64/server/entity"
	"github.com/variant64/server/store"
)

var requestHandler = entity.RequestHandler{}

func handlePostPlayer(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	requestNewPlayer := &entity.RequestNewPlayer{}
	err = json.Unmarshal(body, requestNewPlayer)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}

	player, err := requestHandler.HandleNewPlayer(requestNewPlayer)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	writeEntityResponse(w, player)
}

func handleGetPlayerByID(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	requestGetPlayer := &entity.RequestGetPlayer{ID: id}
	player, err := requestHandler.HandleGetPlayer(requestGetPlayer)
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

	requestNewRoom := &entity.RequestNewRoom{}
	err = json.Unmarshal(body, requestNewRoom)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}

	room, err := requestHandler.HandleNewRoom(requestNewRoom)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	writeEntityResponse(w, room)
}

func handleGetRooms(w http.ResponseWriter, req *http.Request) {
	requestGetRooms := &entity.RequestGetRooms{}
	rooms := requestHandler.HandleGetRooms(requestGetRooms)

	writeEntityListResponse(w, rooms)
}

func handleGetRoomByID(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	requestGetRoom := &entity.RequestGetRoom{ID: id}
	room, err := requestHandler.HandleGetRoom(requestGetRoom)
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, room)
}

func handlePostRoomJoin(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	requestAddPlayer := &entity.RequestRoomAddPlayer{RoomID: id}
	err = json.Unmarshal(body, requestAddPlayer)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}

	room, err := requestHandler.HandleRoomAddPlayer(requestAddPlayer)
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, room)
}

func handlePostRoomLeave(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	requestRemovePlayer := &entity.RequestRoomRemovePlayer{RoomID: id}
	err = json.Unmarshal(body, requestRemovePlayer)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}

	room, err := requestHandler.HandleRoomRemovePlayer(requestRemovePlayer)
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, room)
}

func handlePostRoomStart(w http.ResponseWriter, req *http.Request) {
	id, err := parseIDFromVars(req)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	requestRoomStartGame := &entity.RequestRoomStartGame{RoomID: id}
	err = json.Unmarshal(body, requestRoomStartGame)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}

	room, err := requestHandler.HandleRoomStartGame(requestRoomStartGame)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	writeEntityResponse(w, room)
}

// parseIDFromVars attempts to parse a uuid.UUID from the http.Request path.
func parseIDFromVars(req *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		logger.Error("invalid id", plainError(err))
		return uuid.Nil, errors.Wrap(err, "invalid id")
	}
	return id, nil
}

// writeBadRequestResponse writes an http.StatusBadRequest back to the client.
func writeBadRequestResponse(w http.ResponseWriter, err error) {
	logger.Error(
		"bad request",
		plainError(err),
		zap.Int("status_code", http.StatusInternalServerError),
	)

	w.WriteHeader(http.StatusBadRequest)
	_, err = w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
	if err != nil {
		logger.Error("failed to write response to client", plainError(err))
	}
}

// writeNotFoundResponse writes an http.StatusNotFound back to the client.
func writeNotFoundResponse(w http.ResponseWriter, err error) {
	logger.Error(
		"not found",
		plainError(err),
		zap.Int("status_code", http.StatusInternalServerError),
	)

	w.WriteHeader(http.StatusNotFound)
	_, err = w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
	if err != nil {
		logger.Error("failed to write response to client", plainError(err))
	}
}

// writeServerErrorResponse writes an http.StatusInternalServerError back to the client.
func writeServerErrorResponse(w http.ResponseWriter, err error) {
	logger.Error(
		"internal server error",
		plainError(err),
		zap.Int("status_code", http.StatusInternalServerError),
	)

	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
	if err != nil {
		logger.Error("failed to write response to client", plainError(err))
	}
}

// writeEntityResponse writes the provided entity back to the client.
func writeEntityResponse[T store.Indexable](w http.ResponseWriter, e *entity.Entity[T]) {
	serialized, err := json.Marshal(e.Data)
	if err != nil {
		writeServerErrorResponse(w, err)
	}
	_, err = w.Write(serialized)
	if err != nil {
		writeServerErrorResponse(w, err)
	}
}

// writeEntityListResponse writes the provided entityList back to the client.
func writeEntityListResponse[T store.Indexable](w http.ResponseWriter, e *entity.EntityList[T]) {
	serialized, err := json.Marshal(e.Data)
	if err != nil {
		writeServerErrorResponse(w, err)
	}
	_, err = w.Write(serialized)
	if err != nil {
		writeServerErrorResponse(w, err)
	}
}
