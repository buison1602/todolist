package potal

import "github.com/buison1602/todolist/storage"

type ItemForm struct {
	Item string `json:"item" validate:"omitempty"`
}

func (f *ItemForm) FormCreate() storage.Todo {
	return storage.Todo{
		Item: f.Item,
	}
}

func (f *ItemForm) FormUpdate(todo *storage.Todo) {
	todo.Item = f.Item
}
