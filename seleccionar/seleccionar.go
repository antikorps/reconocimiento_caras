package seleccionar

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func SeleccionarImagenMuestra(ctx context.Context, imagenesReconocimientoRuta string) (string, error) {
	archivo, archivoError := runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Image Files (*.jpg, *.png)",
				Pattern:     "*.jpg;*.png",
			},
		},
	})
	if archivoError != nil {
		return "", archivoError
	}
	if archivo == "" {
		return "", errors.New("proceso cancelado por el usuario")
	}

	nombreImagenReconocimiento := filepath.Base(archivo)
	rutaArchivoImagenReconocimiento := filepath.Join(imagenesReconocimientoRuta, nombreImagenReconocimiento)
	leerImagen, leerImagenError := os.Open(archivo)
	if leerImagenError != nil {
		return "", leerImagenError
	}
	defer leerImagen.Close()

	crearImagen, crearImagenError := os.Create(rutaArchivoImagenReconocimiento)
	if crearImagenError != nil {
		return "", crearImagenError
	}
	defer crearImagen.Close()

	_, copiarImagenError := io.Copy(crearImagen, leerImagen)
	if copiarImagenError != nil {
		return "", copiarImagenError
	}

	return nombreImagenReconocimiento, nil

}
