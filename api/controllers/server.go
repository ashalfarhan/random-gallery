package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	Router *mux.Router
	Port   int
	Addr   string
}

func (s *Server) Init() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 3000
	}
	s.Port = port
	s.Addr = fmt.Sprintf("%s:%d", os.Getenv("HOST"), s.Port)
	s.Router = mux.NewRouter()
	s.InitRoutes()
}

func (s *Server) InitRoutes() {
	s.Router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	}).Methods(http.MethodGet)
	s.Router.HandleFunc("/api/images", CreateImage).Methods(http.MethodPost)
	s.Router.HandleFunc("/api/images", GetAllImages).Methods(http.MethodGet)
	s.Router.HandleFunc("/api/images/{id}", DeleteImage).Methods(http.MethodDelete)
}
