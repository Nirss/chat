package main

import (
	"github.com/Nirss/chat/server"
	"github.com/Nirss/chat/web/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"net/http"
)

func main() {
	app.Route("/", &components.Chat{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Styles: []string{
			"https://unpkg.com/spectre.css/dist/spectre.min.css",
		},
	})
	http.HandleFunc("/chat", server.Chat)
	http.ListenAndServe(":8080", nil)
}
