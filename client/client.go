package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	sender := os.Args[1]
	recepient := os.Args[2]
	msg := os.Args[3]

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, os.Kill)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws/" + sender}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	go func() {
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	time.Sleep(5 * time.Second)

	if recepient != "" {
		c.WriteMessage(websocket.TextMessage, []byte(`{"recepient": "`+recepient+`", "body": "`+msg+`"}`))
	}

	<-done

	c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
