package todo

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

var ErrNotFound = errors.New("todo not found")

type TodoService interface {
	AddTodo(title string) (Todo, error)
	GetTodos() ([]Todo, error)
	RemoveTodo(id string) error
	UpdateTodo(todo TodoUpdate) (Todo, error)
}

type repository interface {
	AddTodo(todo Todo) error
	GetTodos() ([]Todo, error)
	RemoveTodo(id string) error
	UpdateTodo(todo TodoUpdate) (Todo, error)
}

type todoService struct {
	repository repository
}

func NewTodoService(r repository) TodoService {
	return &todoService{
		repository: r,
	}
}

func (t *todoService) AddTodo(title string) (Todo, error) {
	now := time.Now()
	newTodo := Todo{
		ID:        fmt.Sprintf("%05d", rand.Intn(99999)), // TODO: use id generator
		CreatedAt: now,
		UpdatedAt: now,
		Title:     title,
		Completed: false,
	}
	err := t.repository.AddTodo(newTodo)

	return newTodo, err
}

func (t *todoService) UpdateTodo(todo TodoUpdate) (Todo, error) {
	now := time.Now()
	todo.UpdatedAt = &now
	return t.repository.UpdateTodo(todo)
}

func (t *todoService) GetTodos() ([]Todo, error) {
	return t.repository.GetTodos()
}

func (t *todoService) RemoveTodo(id string) error {
	return t.repository.RemoveTodo(id)
}
