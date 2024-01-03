package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todoList/storage"
)

type Server struct {
	db  storage.Storage
	mux *mux.Router
}

func NewWebServer() *Server {
	db := storage.NewStorage()
	var server = Server{
		db:  db,
		mux: mux.NewRouter(), //.StrictSlash(true),
	}
	return &server
}

func (s *Server) Run() {
	s.Route()
	log.Printf("todoList is serving on :%d", 777)
	log.Fatal(http.ListenAndServe(":777", s.mux))
}

func parseJSON(r *http.Request, data interface{}) error {
	var decoder = json.NewDecoder(r.Body)
	var err = decoder.Decode(data)
	defer r.Body.Close()
	return err
}

func response(w http.ResponseWriter, httpStatus int, data interface{}, err error) error {
	w.WriteHeader(httpStatus)
	var encoder = json.NewEncoder(w)
	if err != nil {
		return encoder.Encode(err)
	}
	return encoder.Encode(data)
}
