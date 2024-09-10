package main

import (
	//"errors"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/newmohib/go-lang-bookings-app/internal/config"
	"github.com/newmohib/go-lang-bookings-app/internal/handlers"
	"github.com/newmohib/go-lang-bookings-app/internal/models"
	"github.com/newmohib/go-lang-bookings-app/internal/render"
)

// application portNumber
const portNumber = ":8080"

// initialize app config
// its alos using into middleware or any others into main package
var app config.AppConfig

// initialize sessin manager and its alos using into middleware or any others into main package
var session *scs.SessionManager

// main is the main application function
func main() {
	
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting application on port:", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// run is the main application and it can we will use for testing
func run() error  {

	// what am i going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	// set session
	session = scs.New()
	session.Lifetime = 24 * time.Hour // session will expire after 24 hours
	session.Cookie.Persist = true     // session will expire after the browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// store thsis session into app config
	app.Session = session

	// get the template cache from the carate app config
	tc, err := render.CreateTemplateCache()

	if err != nil {
		fmt.Println("error parsing template:", err)
		log.Fatal("Can not create template cache")
		return err
	}
	app.TamplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo((&app))
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	return nil
	
}
