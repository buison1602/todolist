package web

import (
	"encoding/json"
	"github.com/buison1602/todolist/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Map map[string]interface{}

type Config struct {
	Port int `yaml:"port"`
}

type AppConfig struct {
	Db  storage.Config `yaml:"db"`
	Web Config         `yaml:"web"`
}

type Server struct {
	db  storage.Storage
	mux *mux.Router
	cfg Config
}

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func NewWebServer(c *AppConfig) *Server {
	db := storage.NewStorage(c.Db)
	var server = Server{
		db:  db,
		mux: mux.NewRouter(),
		cfg: c.Web,
	}
	return &server
}

func (s *Server) Run() {
	s.Route()
	log.Printf("todoList is serving on :%d", s.cfg.Port)
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
	res := Response{
		Data: data,
	}
	if err != nil {
		res.Error = err.Error()
	}
	return encoder.Encode(res)
}
