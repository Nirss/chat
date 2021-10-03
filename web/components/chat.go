package components

import (
	"context"
	"encoding/json"
	"github.com/Nirss/chat/common"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"io"
	"log"
	"nhooyr.io/websocket"
	"strings"
)

type Chat struct {
	app.Compo

	name string
	ws   *websocket.Conn
	box  ChatBox
}

func (c *Chat) OnMount(ctx app.Context) {
	var err error
	c.ws, _, err = websocket.Dial(ctx, "ws://localhost:8080/chat", nil)
	if err != nil {
		log.Println("ERROR ", err)
		return
	}
	go c.onMessage()
}

func (c *Chat) printMessage(msg common.Message) {
	c.box.AddMessage(msg)
}

func (c *Chat) onMessage() {
	var (
		msg common.Message
		err error
		nt  websocket.MessageType
		r   io.Reader
	)
	for {
		nt, r, err = c.ws.Reader(context.Background())
		if err != nil {
			log.Println(err)
			return
		}
		if nt == websocket.MessageBinary {
			continue
		}
		err = json.NewDecoder(r).Decode(&msg)
		if err != nil {
			log.Println(err)
			return
		}
		c.printMessage(msg)
	}
}

func (c *Chat) OnKeyUp(ctx app.Context, e app.Event) {
	if e.Get("key").String() == "Enter" && c.name != strings.TrimSpace("") {
		err := c.ws.Write(ctx, websocket.MessageText, []byte(c.name))
		if err != nil {
			log.Println("Write: ", err)
		}
		ctx.JSSrc().Set("value", "")
	}
	c.name = ctx.JSSrc().Get("value").String()
}

func (c *Chat) Render() app.UI {
	return app.Div().Class("container").Body(
		app.Div().Class("columns").Body(
			app.Div().Class("column", "col-4", "col-mx-auto").Body(

				app.Input().
					Class("form-input").
					Type("text").Value(c.name).
					Placeholder("Введите имя").
					AutoFocus(true).
					OnKeyup(c.OnKeyUp),
				&c.box,
			),
		),
	)
}
