package api

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/websocket"
)

// MockHttpServer is a mock http server.
type MockHttpServer struct {
	*httptest.Server
}

// NewMockHttpServer created a new MockHttpServer with injected http.handleFunc.
func NewMockHttpServer(handleFunc http.HandlerFunc) *MockHttpServer {
	server := httptest.NewServer(
		http.HandlerFunc(handleFunc),
	)

	return &MockHttpServer{
		Server: server,
	}
}

// MockWSServer is a mock web server.
type MockWSServer struct {
	*websocket.Conn
}

// NewWebSocketServer creates a mock web server attached to injected httpSeverURL.
func NewWebSocketServer(httpSeverURL string) *MockWSServer {
	serverURL := "ws" + strings.TrimPrefix(httpSeverURL, "http")
	conn, _, _ := websocket.DefaultDialer.Dial(serverURL, nil)

	return &MockWSServer{
		Conn: conn,
	}
}

// NewWebSocketHandleFunc create a http.HandleFuc with http.Upgrade.
func NewWebSocketHandleFunc(handler commandHandler) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				return
			}
			defer conn.Close()

			readAndHandleMessages(conn, handler)
		})
}
