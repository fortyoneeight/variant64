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

type entityWriter[T store.Indexable] interface {
	Write(*entity.Entity[T]) error
}

type entityReader[T store.Indexable] interface {
	Read(*entity.Entity[T]) error
}

func handleNewEntity[T store.Indexable](w http.ResponseWriter, req *http.Request, ew entityWriter[T]) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	err = json.Unmarshal(body, ew)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}

	entity, err := entity.HandleNew[T](ew)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	writeEntityResponse(w, entity)
}

func handleGetEntity[T store.Indexable](w http.ResponseWriter, req *http.Request, ew entityReader[T]) {
	entity, err := entity.HandleGet[T](ew)
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, entity)
}

type actionHandler[T store.Indexable] interface {
	PerformAction() (*entity.Entity[T], error)
	Unmarshal(data []byte) error
}

func handleActionRouteByID[T store.Indexable](w http.ResponseWriter, req *http.Request, handler actionHandler[T]) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to read request body"))
		return
	}

	err = handler.Unmarshal(body)
	if err != nil {
		writeBadRequestResponse(w, errors.Wrap(err, "failed to unmarshal request body"))
		return
	}

	entity, err := handler.PerformAction()
	if err != nil {
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, entity)
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
