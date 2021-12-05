package api

import (
	"log"
	"net/http"
	"time"

	"github.com/ashalfarhan/gallery-api/api/controllers"
)

func Bootstrap() {
	app := controllers.Server{}
	app.Init()
	server := &http.Server{
		Addr:         app.Addr,
		Handler:      app.Router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Printf("Listening on %v\n", app.Addr)
	log.Fatal(server.ListenAndServe())
}
