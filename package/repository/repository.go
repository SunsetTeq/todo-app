package repository

import (
	"github.com/SunsetTeq/firstRestApiProject"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user firstRestApiProject.User) (int, error)
	GetUser(username, password string) (firstRestApiProject.User, error)
}

type TodoList interface {
	Create(userId int, list firstRestApiProject.TodoList) (int, error)
	GetAll(userId int) ([]firstRestApiProject.TodoList, error)
	GetById(userId, listId int) (firstRestApiProject.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input firstRestApiProject.UpdateListInput) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
