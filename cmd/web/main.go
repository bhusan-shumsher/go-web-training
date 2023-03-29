package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhusan-shumsher/go-web-training/pkg/config"
	"github.com/bhusan-shumsher/go-web-training/pkg/handlers"
	"github.com/bhusan-shumsher/go-web-training/pkg/render"
)

const port = ":8000"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache !")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", port))

	_ = http.ListenAndServe(port, nil)
}
