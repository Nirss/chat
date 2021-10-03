package components

import (
	"github.com/Nirss/chat/common"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ChatBox struct {
	app.Compo

	messages []common.Message
}

func (c *ChatBox) Render() app.UI {
	var messages []app.UI

	for _, m := range c.messages {
		messages = append(messages, app.Div().Class("title").Body(
			app.Div().Class("title-content").Body(
				app.P().Class("title-title", "text-bold").Text(m.Nickname)),
			app.P().Class("title-subtitle").Text(m.Message)),
		)
	}

	return app.Div().Class("panel").Body(
		app.Div().Class("panel-header").Body(app.Div().Class("panel-title").Text("Chat")),
		app.Div().Class("panel-body").Body(messages...))
}

func (c *ChatBox) AddMessage(msg common.Message) {
	c.messages = append(c.messages, msg)
	c.Update()
}
