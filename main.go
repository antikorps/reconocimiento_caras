package main

import (
	"embed"
	"log"
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error cargando el archivo .env")
	}
}

func main() {
	app := NuevaApp()
	fs := http.FileServer(http.Dir(filepath.Join(app.EjecutableRuta, "imagenes_reconocimiento")))

	err := wails.Run(&options.App{
		Title:            "reconocimiento_caras",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.iniciar,
		OnBeforeClose:    app.cerrar,
		AssetsHandler:    http.Handler(fs),
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
