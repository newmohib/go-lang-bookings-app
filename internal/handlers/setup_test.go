package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/choi2k/nosurf"
	"github.com/go-chi/chi/v5"
	"github.com/newmohib/go-lang-bookings-app/internal/config"
	"github.com/newmohib/go-lang-bookings-app/internal/models"
	"github.com/newmohib/go-lang-bookings-app/internal/render"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}


// run is the main application and it can we will use for testing
func getRoutes() error {

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
		// return err
	}
	app.TamplateCache = tc
	app.UseCache = false

	repo := NewRepo((&app))
	NewHandlers(repo)

	render.NewTemplate(&app)

	// routes
	// create route with chi
	mux := chi.NewRouter()

	// middleware is used for logging console
	mux.Use(WriteToConsole)
	// set coockie into request as csrf on client
	mux.Use(NoSurf)
	// set session
	mux.Use(SessionLoad)

	// router
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)

	mux.Get("/search-availability", Repo.SearchAvailability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)

	mux.Get("/contact", Repo.Contact)
	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	// static files config
	fileServer := http.FileServer(http.Dir("./static/"))
	// fmt.Println("File Server path: ", fileServer)
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return nil

}



func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})

}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	//app.InProduction is get from main.go due to run the go run cmd/web/*.go

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	// this session is created in main.go as global variable for only main package
	return session.LoadAndSave(next)

}

// crateTemplateCache creates a template cache as a map

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currntly processing", name)
		// we can modify the template into functions here
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		// in tutorial are ./templates/*.layout.tmpl
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}