package procesar

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func (rp *ResultadoProcesamiento) cuadrosDelimitadoresACoordenadas() {
	for _, v := range rp.CuadrosDelimitadores {

		caraIzquierda := v.Left * float64(rp.Ancho)
		caraSuperior := v.Top * float64(rp.Alto)

		caraAncho := v.Width * float64(rp.Ancho)
		caraAlto := v.Height * float64(rp.Alto)

		x1 := int64(caraIzquierda)
		x2 := int64(caraIzquierda + caraAncho)
		y1 := int64(caraSuperior)
		y2 := int64(caraSuperior + caraAlto)

		rp.CoordenadasManuales = append(rp.CoordenadasManuales, []int64{x1, x2, y1, y2})
	}
}

func (rp *ResultadoProcesamiento) pintarDetecciones() error {
	dominioImagen := rp.ImagenDecodificada.Bounds()

	nuevaImagen := image.NewRGBA(dominioImagen)
	for x := dominioImagen.Min.X; x <= dominioImagen.Max.X; x++ {
		for y := dominioImagen.Min.Y; y <= dominioImagen.Max.Y; y++ {
			p := image.Pt(x, y)

			var coincidencia bool
			for _, v := range rp.CoordenadasManuales {
				if p.X >= int(v[0]) && p.X <= int(v[1]) && p.Y == int(v[2]) {
					nuevaImagen.Set(x, y, color.RGBA{0, 255, 0, 255})
					coincidencia = true
				}
				if p.X >= int(v[0]) && p.X <= int(v[1]) && p.Y == int(v[3]) {
					nuevaImagen.Set(x, y, color.RGBA{0, 255, 0, 255})
					coincidencia = true

				}
				if p.Y >= int(v[2]) && p.Y <= int(v[3]) && p.X == int(v[0]) {
					nuevaImagen.Set(x, y, color.RGBA{0, 255, 0, 255})
					coincidencia = true

				}
				if p.Y >= int(v[2]) && p.Y <= int(v[3]) && p.X == int(v[1]) {
					nuevaImagen.Set(x, y, color.RGBA{0, 255, 0, 255})
					coincidencia = true

				}
			}
			if !coincidencia {
				nuevaImagen.Set(x, y, rp.ImagenDecodificada.At(x, y))
			}

		}
	}

	nombreArchivo := filepath.Base(rp.Ruta)
	nuevoNombre := fmt.Sprintf("deteccion_caras_%v", nombreArchivo)
	nuevaRuta := filepath.Join(rp.RutaDestino, nuevoNombre)
	rp.ImagenResultanteNombre = nuevoNombre

	archivoDestino, archivoDestinoError := os.Create(nuevaRuta)
	if archivoDestinoError != nil {
		return archivoDestinoError
	}

	if strings.HasPrefix(nombreArchivo, ".png") {
		crearImagenError := png.Encode(archivoDestino, nuevaImagen)
		if crearImagenError != nil {
			return crearImagenError
		}
	} else {
		crearImagenError := jpeg.Encode(archivoDestino, nuevaImagen, nil)
		if crearImagenError != nil {
			return crearImagenError
		}
	}

	rp.ImagenResultanteNombre = nuevoNombre
	var imagenDecodificada image.Image
	rp.ImagenDecodificada = imagenDecodificada
	return nil

}
