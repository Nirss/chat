package main

import "time"

type Message struct {
	Time     time.Time
	Nickname string
	Message  string
}
