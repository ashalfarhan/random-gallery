package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ashalfarhan/random-gallery/api/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Bootstrap() {
	addr := Init()
	r := InitRoutes()
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	r.Use(cors, JSONContentType)

	server := &http.Server{
		Addr:         addr,
		Handler:      CreateHttpLogger(r),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Printf("Listening on %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func Init() string {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
		log.Printf("Using default port number %d\n", port)
	}

	return fmt.Sprintf("%s:%d", os.Getenv("HOST"), port)
}

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Hello).Methods(http.MethodGet)
	r.HandleFunc("/ping", controllers.Ping).Methods(http.MethodGet)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/images", controllers.CreateImage).Methods(http.MethodPost)
	api.HandleFunc("/images", controllers.GetAllImages).Methods(http.MethodGet)
	api.HandleFunc("/images/{id}", controllers.DeleteImage).Methods(http.MethodDelete)

	return r
}

func CreateHttpLogger(r http.Handler) http.Handler {
	env := os.Getenv("APP_ENV")
	out := os.Stdout
	if env == "prod" {
		return handlers.CombinedLoggingHandler(out, r)
	}
	return handlers.LoggingHandler(out, r)
}

func JSONContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
