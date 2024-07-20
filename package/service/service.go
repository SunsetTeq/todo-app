package service

import (
	"github.com/SunsetTeq/firstRestApiProject"
	"github.com/SunsetTeq/firstRestApiProject/package/repository"
)

type Authorization interface {
	CreateUser(user firstRestApiProject.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list firstRestApiProject.TodoList) (int, error)
	GetAll(userId int) ([]firstRestApiProject.TodoList, error)
	GetById(userId, listId int) (firstRestApiProject.TodoList, error)
	Delete(userId, listId int) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
