package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/variant64/server/entity"
	"github.com/variant64/server/room"
	"github.com/variant64/server/store"
)

var requestHandler = room.RequestHandler{}

type actionHandler[T store.Indexable] interface {
	PerformAction() (*entity.Entity[T], error)
}

func handleActionRoute[T store.Indexable](w http.ResponseWriter, req *http.Request, handler actionHandler[T]) {
	err := parseRequestParameters(req, handler)
	if err != nil {
		writeBadRequestResponse(w, err)
		return
	}

	entity, err := handler.PerformAction()
	if err != nil {
		// TODO: Return a typed error from PerformAction to determine status code.
		writeNotFoundResponse(w, err)
		return
	}

	writeEntityResponse(w, entity)
}

// parseRequestParameters parses the request body and path parameters into the interface.
func parseRequestParameters(req *http.Request, i interface{}) error {
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
				if f.Kind() != reflect.String {
					return data, nil
				}
				if t != reflect.TypeOf(uuid.UUID{}) {
					return data, nil
				}

				return uuid.Parse(data.(string))
			},
		),
		Result: i,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(mux.Vars(req))
	if err != nil {
		return errors.Wrap(err, "failed to decode path parameter")
	}

	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, i)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal request body")
		}
	}

	return nil
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
