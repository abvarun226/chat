package server

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

// Listener method listens to new messages from the user who has established a ws connection with server.
func (s *Server) Listener(u *User) {
	for {
		mt, message, err := u.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[ERROR] failed to read message: %v", err)
			}
			delete(s.Connections, u.Username)
			break
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Printf("[ERROR] failed to unmarshal messsage: %v", err)
			u.conn.WriteMessage(mt, []byte(MessageDecodeError))
			continue
		}

		if msg.Recepient == global {
			s.SendToAll(u.Username, &msg)
			continue
		}

		if _, found := s.Connections[msg.Recepient]; !found {
			log.Printf("[DEBUG] recipient is offline, storing in database")
			continue
		}

		recepientConn := s.Connections[msg.Recepient]
		recepientConn.conn.WriteMessage(mt, []byte(msg.Body))
	}
}

// SendToAll method broadcasts a message to all users connected to the server.
func (s *Server) SendToAll(sender string, msg *Message) {
	for username, user := range s.Connections {
		if sender == username {
			continue
		}
		user.conn.WriteMessage(websocket.TextMessage, []byte(msg.Body))
	}
}

type User struct {
	Username string
	conn     *websocket.Conn
}

type Message struct {
	Recepient string `json:"recepient"`
	Body      string `json:"body"`
}

const (
	global = "GLOBAL"

	MessageDecodeError = `message decoding failed`
)
