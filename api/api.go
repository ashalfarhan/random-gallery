package api

import (
	"log"
	"net/http"
	"time"

	"github.com/ashalfarhan/random-gallery/api/controllers"
	"github.com/rs/cors"
)

func Bootstrap() {
	app := controllers.Server{}
	app.Init()
	c := cors.AllowAll()
	server := &http.Server{
		Addr:         app.Addr,
		Handler:      c.Handler(app.Router),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Printf("Listening on %v\n", app.Addr)
	log.Fatal(server.ListenAndServe())
}
