package main

import (
	"embed"

	"dev.acevedo/backend/log"
	"dev.acevedo/backend/repositories"
	"dev.acevedo/backend/services"

	"dev.acevedo/backend/database"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	//logger init
	log := log.New()

	//database connection
	dbPool, err := database.OpenDb()
	if err != nil {
		log.Error("Database connection failed with ", "err", err.Error())
	}
	defer dbPool.Close()

	taskRepo := repositories.NewTaskRepo(dbPool)
	TaskService := services.NewTaskService(log, taskRepo)
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "ILMS",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			TaskService,
		},
	})

	if err != nil {
		log.Error(err.Error())
	}
}
