package storage

import (
	"bufio"
	"encoding/json"
	"io"
	"slices"
	"todolist-wails/internal/todo"
)

type file interface {
	io.Writer
	io.Seeker
	io.Reader
	Truncate(size int64) error
}

type JSONStorage struct {
	file  file
	todos []todo.Todo
}

func NewJSONStorage(f file) (*JSONStorage, error) {
	j := &JSONStorage{
		file: f,
	}

	if err := j.initializeTodos(); err != nil {
		return nil, err
	}

	return j, nil
}

func (j *JSONStorage) AddTodo(t todo.Todo) error {
	newTodos := append(j.todos, t)

	if err := j.saveTodos(newTodos); err != nil {
		return err
	}

	j.todos = newTodos

	return nil
}

func (j *JSONStorage) RemoveTodo(id string) error {
	newTodos := slices.DeleteFunc(j.todos, func(t todo.Todo) bool {
		return t.ID == id
	})
	if err := j.saveTodos(newTodos); err != nil {
		return err
	}

	j.todos = newTodos

	return nil
}

func (j *JSONStorage) UpdateTodo(u todo.TodoUpdate) (todo.Todo, error) {
	i := slices.IndexFunc(j.todos, func(t todo.Todo) bool {
		return t.ID == u.ID
	})

	if i < 0 {
		return todo.Todo{}, todo.ErrNotFound
	}

	t := j.todos[i]
	t.UpdatedAt = *u.UpdatedAt

	if u.Completed != nil {
		t.Completed = *u.Completed
	}

	if u.Title != nil {
		t.Title = *u.Title
	}

	j.todos[i] = t

	if err := j.saveTodos(j.todos); err != nil {
		return todo.Todo{}, err
	}

	return t, nil
}

func (j *JSONStorage) GetTodos() ([]todo.Todo, error) {
	return j.todos, nil
}

func (j *JSONStorage) saveTodos(todos []todo.Todo) error {
	data, mErr := json.Marshal(todos)
	if mErr != nil {
		return mErr
	}

	if err := j.file.Truncate(0); err != nil {
		return err
	}

	if _, err := j.file.Seek(0, 0); err != nil {
		return err
	}

	if _, err := j.file.Write(data); err != nil {
		return err
	}

	return nil
}

func (j *JSONStorage) initializeTodos() error {
	scanner := bufio.NewScanner(j.file)
	var data []byte

	for scanner.Scan() {
		data = append(data, scanner.Bytes()...)
	}

	todos := make([]todo.Todo, 0)

	if len(data) > 0 {
		if err := json.Unmarshal(data, &todos); err != nil {
			return err
		}
	}

	j.todos = todos

	return nil
}
