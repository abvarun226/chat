package main

import (
	"log"
	"net/http"

	"github.com/abvarun226/chat/server"
	"github.com/go-chi/chi"
)

const (
	serverAddr = ":8080"
)

func main() {
	r := chi.NewRouter()

	s := server.New()

	r.Get("/status", Status)
	r.Get("/ws/{user}", s.WebsocketHandler)

	log.Printf("[INFO] starting server on %s", serverAddr)
	http.ListenAndServe(serverAddr, r)
}

// Status handler is used to check the health of the server.
func Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
