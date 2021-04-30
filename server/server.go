package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

// Server is our websocket server.
type Server struct {
	sync.Mutex
	Connections map[string]*User
	Upgrader    websocket.Upgrader
}

// New creates a new websocket server.
func New() *Server {
	return &Server{
		Connections: make(map[string]*User),
		Upgrader:    websocket.Upgrader{},
	}
}

// WebsocketHandler handles a websocket conn initiation and starts a listener.
func (s *Server) WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")

	wsConnection, errUpgrade := s.Upgrader.Upgrade(w, r, nil)
	if errUpgrade != nil {
		log.Printf("[ERROR] failed to upgrade ws connection: %v", errUpgrade)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to establish ws connection"))
		return
	}
	defer wsConnection.Close()

	u := &User{Username: user, conn: wsConnection}

	if _, found := s.Connections[user]; !found {
		s.Lock()
		s.Connections[user] = u
		s.Unlock()
	}

	s.Listener(u)
}
