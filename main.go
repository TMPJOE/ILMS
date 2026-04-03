package main

import (
	"embed"
	"os"
	"path/filepath"
	"strings"

	"dev.acevedo/backend"
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
	// Get the directory where the binary is located
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execDir := filepath.Dir(execPath)

	// Use current working directory if we're in dev mode (binary is in a temp location)
	// Check if execDir contains "wails" or is a temp directory
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	// If running from Wails dev mode, use the current working directory instead
	if !filepath.IsAbs(execDir) || strings.Contains(execDir, "wails") || strings.Contains(execDir, "tmp") {
		execDir = cwd
	}

	// Create directories relative to the determined directory
	logsDir := filepath.Join(execDir, "logs")
	dbDir := filepath.Join(execDir, "database")

	err = os.MkdirAll(logsDir, 0755)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(dbDir, 0755)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(filepath.Join(logsDir, "tasks.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	//logger init
	log := log.New(logFile)

	//database connection
	dbPath := filepath.Join(dbDir, "tasks.db")
	dbPool, err := database.OpenDb(backend.DbFiles, log, dbPath)
	if err != nil {
		log.Error("Database connection failed", "error", err.Error())
		panic(err)
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
		log.Error("Error creating app", "err", err.Error())
		panic(err)
	}
}
