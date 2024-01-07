package web

func (s *Server) Route() {
	s.mux.HandleFunc("/api/item", s.createItem).Methods("POST")
	s.mux.HandleFunc("/api/items", s.getListItem).Methods("GET")
	s.mux.HandleFunc("/api/item/{id:[0-9]+}", s.updateItem).Methods("PUT")
	s.mux.HandleFunc("/api/item/{id:[0-9]+}", s.deleteItem).Methods("DELETE")
	s.mux.HandleFunc("/api/auth/register", s.Register).Methods("POST")
	s.mux.HandleFunc("/api/auth/login", s.Login).Methods("POST")
}
