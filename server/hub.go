package server

import (
	"github.com/Nirss/chat/common"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type Hub struct {
	clients []*Client
}

func (h *Hub) ConnectClient(conn *websocket.Conn, nickname string) {
	c := NewClient(conn, nickname)
	h.clients = append(h.clients, c)
	go h.read(c)
	go c.write()
}

func (h *Hub) read(c *Client) {
	for {
		mt, data, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("client %s read error: %+v", c.nickname, err)
			return
		}
		if mt == websocket.BinaryMessage {
			continue
		}
		msg := common.Message{
			Time:     time.Now(),
			Message:  string(data),
			Nickname: c.nickname,
		}
		h.Broadcast(msg)
	}
}

func (h *Hub) Broadcast(msg common.Message) {
	for i := range h.clients {
		h.clients[i].SendMessage(msg)
	}
}
