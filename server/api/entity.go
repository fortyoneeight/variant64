package api

import (
	"encoding/json"
	"net/http"

	"github.com/variant64/server/entity"
	"github.com/variant64/server/store"
)

func handleReadEntity[T store.Indexable] (
	w http.ResponseWriter, req *http.Request, entityReq entity.EntityReadRequest[T],
) {
	entity := &entity.Entity[T]{}
	err := entityReq.Read(entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = entity.Load()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serialized, _ := json.Marshal(entity.Data)

	w.Write(serialized)
}

func handleReadEntities[T store.Indexable] (
	w http.ResponseWriter, req *http.Request, entityReq entity.EntityListReadRequest[T],
) {
	entity := &entity.EntityList[T]{}
	err := entityReq.Read(entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entity.Load()

	serialized, _ := json.Marshal(entity.Data)

	w.Write(serialized)
}

func handleWriteEntity[T store.Indexable] (
	w http.ResponseWriter, req *http.Request, entityReq entity.EntityWriteRequest[T],
) {
	err := json.NewDecoder(req.Body).Decode(entityReq)
	if err != nil {
		w.Write([]byte(InvalidBodyResponse))
		return
	}

	entity := &entity.Entity[T]{}
	err = entityReq.Write(entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	entity.Store()

	serialized, _ := json.Marshal(entity.Data)

	w.Write(serialized)
}
