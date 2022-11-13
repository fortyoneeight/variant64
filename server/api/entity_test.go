package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/variant64/server/entity"
)

func TestHandleReadEntity(t *testing.T) {
	entityID := uuid.New()

	testcases := []struct {
		name                 string
		readRequest          entity.EntityReadRequest[entity.MockEntity]
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Valid request, status OK.",
			readRequest:          &entity.MockEntityReadRequestSuccess{ID: entityID},
			expectedStatusCode:   200,
			expectedResponseBody: fmt.Sprintf("{\"id\":\"%s\",\"field\":\"\"}", entityID),
		},
		{
			name:                 "Invalid Read, status BadRequest.",
			readRequest:          &entity.MockEntityReadRequestFailed{ID: entityID},
			expectedStatusCode:   400,
			expectedResponseBody: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockResponseWriter := &mockResponseWriter{
				statusCode: 200,
			}
			mockRequest := &http.Request{}

			handleReadEntity(mockResponseWriter, mockRequest, tc.readRequest)

			assert.Equal(t, tc.expectedStatusCode, mockResponseWriter.statusCode)
			assert.Equal(t, []byte(tc.expectedResponseBody), mockResponseWriter.bytes)
		})
	}
}

func TestHandleReadEntities(t *testing.T) {
	entityID := uuid.New()

	testcases := []struct {
		name                 string
		readRequest          entity.EntityListReadRequest[entity.MockEntity]
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Valid request, status OK.",
			readRequest:          &entity.MockEntityListReadRequestSuccess{ID: entityID},
			expectedStatusCode:   200,
			expectedResponseBody: fmt.Sprintf("[{\"id\":\"%s\",\"field\":\"\"}]", entityID),
		},
		{
			name:                 "Invalid Read, status BadRequest.",
			readRequest:          &entity.MockEntityListReadRequestFailed{ID: entityID},
			expectedStatusCode:   400,
			expectedResponseBody: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockResponseWriter := &mockResponseWriter{
				statusCode: 200,
			}
			mockRequest := &http.Request{}

			handleReadEntities(mockResponseWriter, mockRequest, tc.readRequest)

			assert.Equal(t, tc.expectedStatusCode, mockResponseWriter.statusCode)
			assert.Equal(t, []byte(tc.expectedResponseBody), mockResponseWriter.bytes)
		})
	}
}

func TestHandleWriteEntity(t *testing.T) {
	entityID := uuid.New()

	testcases := []struct {
		name                 string
		writeRequest         entity.EntityWriteRequest[entity.MockEntity]
		requestBody          string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Valid request, status OK.",
			writeRequest:         &entity.MockEntityWriteRequestSuccess{},
			requestBody:          fmt.Sprintf("{\"id\":\"%s\",\"field\":\"val\"}", entityID),
			expectedStatusCode:   200,
			expectedResponseBody: fmt.Sprintf("{\"id\":\"%s\",\"field\":\"val\"}", entityID),
		},
		{
			name:                 "Invalid Write, status BadRequest.",
			writeRequest:         &entity.MockEntityWriteRequestFailed{},
			requestBody:          fmt.Sprintf("{\"id\":\"%s\",\"field\":\"val\"}", entityID),
			expectedStatusCode:   400,
			expectedResponseBody: "",
		},
		{
			name:                 "Invalid request, status BadRequest.",
			writeRequest:         &entity.MockEntityWriteRequestSuccess{},
			requestBody:          "",
			expectedStatusCode:   400,
			expectedResponseBody: invalidBodyResponse,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockResponseWriter := &mockResponseWriter{
				statusCode: 200,
			}
			mockRequest := &http.Request{Body: io.NopCloser(strings.NewReader(tc.requestBody))}

			handleWriteEntity(mockResponseWriter, mockRequest, tc.writeRequest)

			assert.Equal(t, tc.expectedStatusCode, mockResponseWriter.statusCode)
			assert.Equal(t, []byte(tc.expectedResponseBody), mockResponseWriter.bytes)
		})
	}
}

type mockResponseWriter struct {
	statusCode int
	bytes      []byte
}

func (m *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *mockResponseWriter) Write(bytes []byte) (int, error) {
	m.bytes = bytes
	return 0, nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}
