package api

import (
	"encoding/json"
	"net/http"

	"github.com/variant64/server/entity"
	"github.com/variant64/server/store"
)

func handleReadEntity[T store.Indexable](
	// handleReadEntity performs a generic entity.EntityReadRequest
	// the result of the read is serialized and returned via the http.Request.
	w http.ResponseWriter, req *http.Request, entityReq entity.EntityReadRequest[T],
) {
	entity := &entity.Entity[T]{}
	err := entityReq.Read(entity)
	if err != nil {
		w.Write([]byte{})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = entity.Load()
	if err != nil {
		w.Write([]byte{})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serialized, _ := json.Marshal(entity.Data)

	w.Write(serialized)
}

// handleReadEntities performs a generic entity.EntityListReadRequest
// the results of the read are serialized and returned via the http.Request.
func handleReadEntities[T store.Indexable](
	w http.ResponseWriter, req *http.Request, entityReq entity.EntityListReadRequest[T],
) {
	entity := &entity.EntityList[T]{}
	err := entityReq.Read(entity)
	if err != nil {
		w.Write([]byte{})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entity.Load()

	serialized, _ := json.Marshal(entity.Data)

	w.Write(serialized)
}

// handleReadEntitiy performs a generic entity.EntityWriteRequest
// the resulting entity is serialized and returned via the http.Request.
func handleWriteEntity[T store.Indexable](
	w http.ResponseWriter, req *http.Request, entityReq entity.EntityWriteRequest[T],
) {
	err := json.NewDecoder(req.Body).Decode(entityReq)
	if err != nil {
		w.Write([]byte(invalidBodyResponse))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	entity := &entity.Entity[T]{}
	err = entityReq.Write(entity)
	if err != nil {
		w.Write([]byte{})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entity.Store()

	serialized, _ := json.Marshal(entity.Data)

	w.Write(serialized)
}
