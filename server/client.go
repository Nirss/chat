package server

import (
	"github.com/Nirss/chat/common"
	"github.com/gorilla/websocket"
)

type Client struct {
	nickname string
	conn     *websocket.Conn
	ch       chan common.Message
}

func NewClient(conn *websocket.Conn, nickname string) *Client {
	return &Client{
		conn:     conn,
		nickname: nickname,
		ch:       make(chan common.Message),
	}
}

func (c *Client) SendMessage(msg common.Message) {
	c.ch <- msg
}

func (c *Client) write() {
	for {
		msg := <-c.ch
		c.conn.WriteJSON(msg)
	}
}
