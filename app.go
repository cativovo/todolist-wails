package main

import (
	"context"
	"todolist-wails/internal/todo"
)

// App struct
type App struct {
	ctx         context.Context
	todoService todo.TodoService
}

// NewApp creates a new App application struct
func NewApp(t todo.TodoService) *App {
	return &App{
		todoService: t,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetTodos() ([]todo.Todo, error) {
	return a.todoService.GetTodos()
}

func (a *App) AddTodo(title string) (todo.Todo, error) {
	return a.todoService.AddTodo(title)
}

func (a *App) RemoveTodo(id string) error {
	return a.todoService.RemoveTodo(id)
}
