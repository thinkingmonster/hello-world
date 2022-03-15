package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/thinkingmonster/hello-world/pkg/config"
	"github.com/thinkingmonster/hello-world/pkg/handlers"
	"github.com/thinkingmonster/hello-world/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

// main is the main function
var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	fmt.Println(app.TemplateCache)
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
