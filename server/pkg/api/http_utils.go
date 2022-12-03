package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/variant64/server/pkg/errortypes"
	"go.uber.org/zap"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type actionHandler[T any] interface {
	PerformAction() (T, errortypes.TypedError)
}

func handleActionRoute[T any](w http.ResponseWriter, req *http.Request, handler actionHandler[T]) {
	err := parseRequestParameters(req, handler)
	if err != nil {
		writeStatusCode(w, http.StatusBadRequest, err)
		return
	}

	entity, actionErr := handler.PerformAction()
	if actionErr != nil {
		switch actionErr.GetType() {
		case errortypes.NotFound:
			writeStatusCode(w, http.StatusNotFound, actionErr)
			return
		case errortypes.BadRequest:
			writeStatusCode(w, http.StatusBadRequest, actionErr)
			return
		case errortypes.InternalError:
			writeStatusCode(w, http.StatusInternalServerError, actionErr)
			return
		}
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

	if req.Method != "GET" {
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

// writeStatusCode writes a status code back to the client.
func writeStatusCode(w http.ResponseWriter, statusCode int, err error) {
	logger.Error(
		err.Error(),
		zap.Int("status_code", statusCode),
	)

	w.WriteHeader(statusCode)
	_, err = w.Write([]byte(fmt.Sprintf("{\"error\":\"%s\"}", err.Error())))
	if err != nil {
		logger.Error("failed to write response to client", zap.Error(err))
	}
}

// writeEntityResponse writes the provided entity back to the client.
func writeEntityResponse[T any](w http.ResponseWriter, e T) {
	serialized, err := json.Marshal(e)
	if err != nil {
		writeStatusCode(w, http.StatusInternalServerError, err)
	}
	_, err = w.Write(serialized)
	if err != nil {
		writeStatusCode(w, http.StatusInternalServerError, err)
	}
}
