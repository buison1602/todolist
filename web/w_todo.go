package web

import (
	"github.com/buison1602/todolist/helper"
	"github.com/buison1602/todolist/storage"
	"github.com/buison1602/todolist/web/potal"
	"github.com/gorilla/mux"
	"net/http"
)

func (s *Server) createItem(w http.ResponseWriter, r *http.Request) {
	var f potal.ItemForm
	err := parseJSON(r, &f)
	if err != nil {
		response(w, http.StatusBadRequest, nil, err)
		return
	}
	todo := f.FormCreate()
	err = s.db.Create(&todo)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}
	response(w, http.StatusCreated, todo, nil)
}

func (s *Server) getListItem(w http.ResponseWriter, r *http.Request) {
	var todos []storage.Todo
	err := s.db.FindQuery(&todos)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
	}
	response(w, http.StatusOK, todos, nil)
}

func (s *Server) updateItem(w http.ResponseWriter, r *http.Request) {
	var f potal.ItemForm
	err := parseJSON(r, &f)
	if err != nil {
		response(w, http.StatusBadRequest, nil, err)
		return
	}
	vars := mux.Vars(r)
	id := helper.ToDbId(vars["id"])
	var todo storage.Todo
	err = s.db.FirstById(&todo, id)
	if err != nil {
		response(w, http.StatusBadRequest, nil, err)
		return
	}
	f.FormUpdate(&todo)
	err = s.db.Save(&todo)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}
	response(w, http.StatusCreated, todo, nil)
}

func (s *Server) deleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := helper.ToDbId(vars["id"])
	var todo storage.Todo
	err := s.db.FirstById(&todo, id)
	if err != nil {
		response(w, http.StatusBadRequest, nil, err)
		return
	}
	err = s.db.Delete(&todo)
	if err != nil {
		response(w, http.StatusInternalServerError, nil, err)
		return
	}
	response(w, http.StatusOK, todo, nil)
}
