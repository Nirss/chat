package main

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	nickname string
	conn     *websocket.Conn
	ch       chan Message
}

func NewClient(conn *websocket.Conn, nickname string) *Client {
	return &Client{
		conn:     conn,
		nickname: nickname,
		ch:       make(chan Message),
	}
}

func (c *Client) SendMessage(msg Message) {
	c.ch <- msg
}

func (c *Client) write() {
	for {
		msg := <-c.ch
		c.conn.WriteJSON(msg)
	}
}
