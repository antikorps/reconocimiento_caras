package explorar

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reconocimiento_caras/utilidades"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type ImagenesExploradas struct {
	Validas   []AnalisisExploracion `json:"validas"`
	Invalidas []AnalisisExploracion `json:"invalidas"`
}

type AnalisisExploracion struct {
	Ruta  string `json:"ruta"`
	Error string `json:"error"`
}

func ExplorarImagenes(ctx context.Context) (ImagenesExploradas, error) {
	var imagenesExploradas ImagenesExploradas

	directorio, directorioError := runtime.OpenDirectoryDialog(ctx, runtime.OpenDialogOptions{})
	if directorioError != nil {
		return imagenesExploradas, directorioError
	}

	if directorio == "" {
		return imagenesExploradas, errors.New("proceso cancelado por el usuario")
	}

	filepath.Walk(directorio, func(ruta string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		var analisisExploracion AnalisisExploracion
		analisisExploracion.Ruta = ruta

		if err != nil {
			analisisExploracion.Error = err.Error()
			imagenesExploradas.Invalidas = append(imagenesExploradas.Invalidas, analisisExploracion)
			return nil
		}
		extensionArchivo := filepath.Ext(ruta)
		if !utilidades.ExtensionPermitida(extensionArchivo) {
			analisisExploracion.Error = "extensión de archivo no válido"
			imagenesExploradas.Invalidas = append(imagenesExploradas.Invalidas, analisisExploracion)
			return nil
		}

		var maxBytes int64
		if os.Getenv("IMAGEN_MAX_BYTES") == "" {
			maxBytes = 5000000
		} else {
			convertirBytes, convertirBytesError := strconv.Atoi(os.Getenv("IMAGEN_MAX_BYTES"))
			if convertirBytesError != nil {
				log.Println(convertirBytesError)
				maxBytes = 5000000
			}
			maxBytes = int64(convertirBytes)
		}
		if info.Size() > maxBytes {
			analisisExploracion.Error = fmt.Sprintf("tamaño de archivo superior %d al permitido (%d)", info.Size(), maxBytes)
			imagenesExploradas.Invalidas = append(imagenesExploradas.Invalidas, analisisExploracion)
			return nil
		}

		imagenesExploradas.Validas = append(imagenesExploradas.Validas, analisisExploracion)
		return nil
	})
	return imagenesExploradas, nil
}
