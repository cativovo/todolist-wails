package storage_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
	"todolist-wails/internal/storage"
	"todolist-wails/internal/todo"

	"github.com/stretchr/testify/assert"
)

type fileMock struct {
	bytes.Buffer
	offset int64
	whence int
	size   int64
}

func (f *fileMock) Seek(offset int64, whence int) (int64, error) {
	f.offset = offset
	f.whence = whence
	return 0, nil
}

func (f *fileMock) Truncate(size int64) error {
	f.size = size
	return nil
}

func newJSONStorage(t *testing.T, fileContent []byte) (storage.JSONStorage, *fileMock) {
	t.Helper()

	file := fileMock{
		size:   -1,
		offset: -1,
		whence: -1,
	}
	file.Buffer = *bytes.NewBuffer(fileContent)
	jsonStorage, err := storage.NewJSONStorage(&file)
	assert.Nil(t, err)
	return *jsonStorage, &file
}

func assertSaveToFile(t *testing.T, f *fileMock, todos []todo.Todo) {
	t.Helper()

	assert.Equal(t, 0, f.whence)
	assert.Equal(t, int64(0), f.offset)
	assert.Equal(t, int64(0), f.size)

	actual, err := json.Marshal(&todos)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, f.Buffer.Bytes(), actual)
}

func TestJSONStorage(t *testing.T) {
	testTodos := []todo.Todo{
		{
			CreatedAt: time.Date(2024, 9, 1, 9, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 9, 1, 9, 0, 0, 0, time.UTC),
			ID:        "1",
			Title:     "Buy groceries",
			Completed: false,
		},
		{
			CreatedAt: time.Date(2024, 9, 2, 10, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 9, 2, 10, 0, 0, 0, time.UTC),
			ID:        "2",
			Title:     "Read a book",
			Completed: false,
		},
		{
			CreatedAt: time.Date(2024, 9, 3, 11, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 9, 3, 11, 0, 0, 0, time.UTC),
			ID:        "3",
			Title:     "Exercise",
			Completed: true,
		},
		{
			CreatedAt: time.Date(2024, 9, 4, 12, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 9, 4, 12, 0, 0, 0, time.UTC),
			ID:        "4",
			Title:     "Clean the house",
			Completed: false,
		},
		{
			CreatedAt: time.Date(2024, 9, 5, 13, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 9, 5, 13, 0, 0, 0, time.UTC),
			ID:        "5",
			Title:     "Write code",
			Completed: true,
		},
	}

	fileContent, mErr := json.Marshal(&testTodos)
	if mErr != nil {
		t.Fatal(mErr)
	}

	t.Run("get todos", func(t *testing.T) {
		jsonStorage, _ := newJSONStorage(t, fileContent)

		todos, err := jsonStorage.GetTodos()
		assert.Nil(t, err)
		assert.Equal(t, testTodos, todos)
	})

	t.Run("add todo", func(t *testing.T) {
		expectedTodo := todo.Todo{
			CreatedAt: time.Date(2024, 9, 5, 13, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2024, 9, 5, 13, 0, 0, 0, time.UTC),
			ID:        "69",
			Title:     "new todo",
			Completed: false,
		}
		expectedTodos := append(testTodos, expectedTodo)
		jsonStorage, f := newJSONStorage(t, fileContent)

		err := jsonStorage.AddTodo(expectedTodo)
		assert.Nil(t, err)
		assertSaveToFile(t, f, expectedTodos)

		todos, err := jsonStorage.GetTodos()
		assert.Nil(t, err)
		assert.Equal(t, expectedTodos, todos)
	})

	t.Run("remove todo", func(t *testing.T) {
		expectedTodos := make([]todo.Todo, len(testTodos))
		copy(expectedTodos, testTodos)
		expectedTodos = append(expectedTodos[:2], expectedTodos[3:]...)

		fileContentCopy := make([]byte, len(fileContent))
		copy(fileContentCopy, fileContent)
		jsonStorage, f := newJSONStorage(t, fileContentCopy)

		err := jsonStorage.RemoveTodo("3")
		assert.Nil(t, err)
		assertSaveToFile(t, f, expectedTodos)

		todos, err := jsonStorage.GetTodos()
		assert.Nil(t, err)
		assert.Equal(t, expectedTodos, todos)
	})

	t.Run("update todo", func(t *testing.T) {
		expectedTodos := make([]todo.Todo, len(testTodos))
		copy(expectedTodos, testTodos)

		now := time.Now()
		expectedTitle := "new title"
		expectedCompleted := true
		index := 1
		expectedTodos[index].Title = expectedTitle
		expectedTodos[index].Completed = expectedCompleted
		expectedTodos[index].UpdatedAt = now
		jsonStorage, f := newJSONStorage(t, fileContent)

		actualTodo, err := jsonStorage.UpdateTodo(todo.TodoUpdate{
			ID:        "2",
			Title:     &expectedTitle,
			Completed: &expectedCompleted,
			UpdatedAt: now,
		})
		assert.Equal(t, expectedTodos[index], actualTodo)
		assert.Nil(t, err)
		assertSaveToFile(t, f, expectedTodos)

		todos, err := jsonStorage.GetTodos()
		assert.Nil(t, err)
		assert.Equal(t, expectedTodos, todos)
	})
}
