package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/sanyog2491/bookings/package/config"
	"github.com/sanyog2491/bookings/package/handlers"
	"github.com/sanyog2491/bookings/package/render"
)

const portnum3 = ":8020"

var app config.Appconfig

var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	//creating the session

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.Createtemplatecache()

	if err != nil {
		log.Println("error while creating template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println("application is activated in port num", portnum3)

	srv := &http.Server{
		Addr:    portnum3,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
