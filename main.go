package main

import (
	"embed"

	"dev.acevedo/backend/features"
	"dev.acevedo/backend/log"

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
	// Create an instance of the app structure
	app := features.NewApp()

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
		},
	})

	if err != nil {
		log.Error(err.Error())
	}
}
