package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"reconocimiento_caras/explorar"
	"reconocimiento_caras/procesar"
	"reconocimiento_caras/seleccionar"
)

// App struct
type App struct {
	Ctx                        context.Context
	EjecutableRuta             string
	ImagenesReconocimientoRuta string
}

// NewApp creates a new App application struct
func NuevaApp() *App {
	ejecutable, ejecutableError := os.Executable()
	if ejecutableError != nil {
		log.Fatalln(ejecutableError)
	}
	rutaEjecutable := filepath.Dir(ejecutable)

	rutaImagenesReconocimiento := filepath.Join(rutaEjecutable, "imagenes_reconocimiento")
	os.Remove(filepath.Join(rutaEjecutable, "imagenes_reconocimiento"))
	os.Mkdir(rutaImagenesReconocimiento, 0777)

	return &App{
		EjecutableRuta:             rutaEjecutable,
		ImagenesReconocimientoRuta: rutaImagenesReconocimiento,
	}
}

func (a *App) iniciar(ctx context.Context) {
	a.Ctx = ctx
}

func (a *App) cerrar(ctx context.Context) (prevent bool) {
	os.RemoveAll(a.ImagenesReconocimientoRuta)
	return false
}

func (a *App) DevolverImagenReconocimiento() (string, error) {
	return seleccionar.SeleccionarImagenMuestra(a.Ctx, a.ImagenesReconocimientoRuta)
}

func (a *App) DevolverImagenesExploradas() (explorar.ImagenesExploradas, error) {
	return explorar.ExplorarImagenes(a.Ctx)
}

func (a *App) DevolverResultadoProcesamiento(nombreImagenMuestra string, imagenesAnalizar []string) procesar.RespuestaProcesamiento {
	return procesar.IniciarProcesamiento(filepath.Join(a.ImagenesReconocimientoRuta, nombreImagenMuestra), imagenesAnalizar)
}
