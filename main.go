package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Text string
}

func chat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade: ", err)
		return
	}
	defer c.Close()

	var m Message
	for {
		c.ReadJSON(&m)
		if err != nil {
			println("ReadJson:", err)
			continue
		}
		err = c.WriteJSON(m)

		if err != nil {
			println("ReadJson:", err)
			continue
		}
	}
}

func main() {
	http.HandleFunc("/chat", chat)
	http.ListenAndServe(":8080", nil)
}
