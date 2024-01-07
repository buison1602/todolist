package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Dns string `yaml:"dns"`
}

// represent for the connection to postgresql
type psql struct {
	db *gorm.DB
}

func NewStorage(c Config) Storage {
	db, err := gorm.Open(postgres.Open(c.Dns), &gorm.Config{})
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
	return conn.AutoMigrate(&Todo{}, &User{})
}

type Storage interface {
	Create(obj interface{}) error
	Save(obj interface{}) error
	Delete(obj interface{}) error
	FindQuery(todos *[]Todo) error
	FirstById(todo *Todo, id int) error
	CheckDuplicate(user *User) error
	FirstByUserName(userName string) (*User, error)
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

func (p *psql) CheckDuplicate(user *User) error {
	var oldUser User
	err := p.db.Where("email = ?", user.Email).First(&oldUser).Error
	if err == nil {
		return fmt.Errorf("email is exist")
	}
	err = p.db.Where("user_name = ?", user.UserName).First(&oldUser).Error
	if err == nil {
		return fmt.Errorf("username is exist")
	}
	return nil
}

func (p *psql) FirstByUserName(userName string) (*User, error) {
	var oldUser User
	err := p.db.Where("user_name = ?", userName).First(&oldUser).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}
	return &oldUser, nil
}
