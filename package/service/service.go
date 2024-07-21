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
	Update(userId, listId int, input firstRestApiProject.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item firstRestApiProject.TodoItem) (int, error)
	GetAll(userId, listId int) ([]firstRestApiProject.TodoItem, error)
	GetById(userId, itemId int) (firstRestApiProject.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input firstRestApiProject.UpdateItemInput) error
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
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
