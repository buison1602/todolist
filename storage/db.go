package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// represent for the connection to postgresql
type psql struct {
	db *gorm.DB
}

func NewStorage() Storage {
	connectionStr := "user=yourtodolist password=password dbname=todolist port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migrate(db)
	var p = &psql{
		db: db,
	}
	return p
}

// create table in Database
func migrate(conn *gorm.DB) error {
	return conn.AutoMigrate(&Todo{})
}

type Storage interface {
	Create(obj interface{}) error
	Save(obj interface{}) error
	Delete(obj interface{}) error
	FindQuery(todos *[]Todo) error
	FirstById(todo *Todo, id int) error
}

func (p *psql) Create(obj interface{}) error {
	return p.db.Create(obj).Error
}

func (p *psql) Save(obj interface{}) error {
	return p.db.Save(obj).Error
}

func (p *psql) Delete(obj interface{}) error {
	return p.db.Delete(obj).Error
}

func (p *psql) FindQuery(todos *[]Todo) error {
	err := p.db.Find(todos).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	return err
}

func (p *psql) FirstById(todo *Todo, id int) error {
	err := p.db.First(&todo, "id = ?", id).Error
	if err == gorm.ErrRecordNotFound {
		return err
	}
	return nil
}
