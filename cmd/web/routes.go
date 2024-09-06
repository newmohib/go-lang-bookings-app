package main

import (
	"net/http"

	// "github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/newmohib/go-lang-bookings-app/internal/config"
	"github.com/newmohib/go-lang-bookings-app/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// routs create with pat
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// create route with chi
	mux := chi.NewRouter()

	// middleware

	mux.Use(middleware.Recoverer)
	// middleware is used for logging console
	mux.Use(WriteToConsole)
	// set coockie into request as csrf on client
	mux.Use(NoSurf)
	// set session
	mux.Use(SessionLoad)

	// router
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.Reservation)

	// static files config
	fileServer := http.FileServer(http.Dir("./static/"))
	// fmt.Println("File Server path: ", fileServer)
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
