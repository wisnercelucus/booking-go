package main

import (
	"log"
	"net/http"

	"github.com/wisnercelucus/udemy-booking/pkg/config"
	"github.com/wisnercelucus/udemy-booking/pkg/handlers"
	"github.com/wisnercelucus/udemy-booking/pkg/render"
)

const PortNumber string = ":8001"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create the template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	http.ListenAndServe(PortNumber, nil)
}
