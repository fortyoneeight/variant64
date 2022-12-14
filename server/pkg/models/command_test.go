package models_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/pkg/models"
)

func TestSubscriberWriteMessage(t *testing.T) {
	mocksEventWriter := models.NewMockEventWriter()
	subscriber := models.NewMessageSubscriber[models.UpdateMessage[string]]("test", mocksEventWriter)
	err := subscriber.OnMessage(
		models.UpdateMessage[string]{
			Channel: "test",
			Data:    "test update",
			Type:    0,
		},
	)
	assert.Nil(t, err)

	var schema = &models.UpdateMessage[string]{}
	err = json.Unmarshal([]byte(mocksEventWriter.LastMessage()), schema)
	assert.Nil(t, err)

	assert.Equal(t, "test", schema.Channel)
	assert.Equal(t, "test update", schema.Data)
	assert.Equal(t, "none", schema.Type.String())
}
