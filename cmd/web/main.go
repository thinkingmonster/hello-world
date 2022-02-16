package main

import (
	"fmt"
	"github.com/thinkingmonster/hello-world/pkg/config"
	"github.com/thinkingmonster/hello-world/pkg/handlers"
	"github.com/thinkingmonster/hello-world/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("can not create template cache")
	}

	app.TemplateCache = tc
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
