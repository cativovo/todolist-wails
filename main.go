package main

import (
	"embed"
	"os"
	"todolist-wails/internal/storage"
	"todolist-wails/internal/todo"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	file, fErr := os.OpenFile("./todos.json", os.O_CREATE|os.O_RDWR, 0644)
	if fErr != nil {
		panic(fErr)
	}
	defer file.Close()

	jsonStorage, jErr := storage.NewJSONStorage(file)
	if jErr != nil {
		panic(jErr)
	}
	todoService := todo.NewTodoService(jsonStorage)

	// Create an instance of the app structure
	app := NewApp(todoService)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "todolist-wails",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
