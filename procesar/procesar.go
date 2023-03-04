package procesar

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func conectarRekognition() *rekognition.Rekognition {
	sess := session.New(&aws.Config{
		Region: aws.String(os.Getenv("REGION"))},
	)
	return rekognition.New(sess)
}

func convertirRekognitionImage(rutaImagen string) (rekognition.Image, error) {
	var r rekognition.Image
	archivo, archivoError := ioutil.ReadFile(rutaImagen)
	if archivoError != nil {
		return r, archivoError
	}
	r.Bytes = archivo
	return r, nil
}

func obtenerInformacionImagen(rutaImagen string) (int, int, image.Image, error) {
	var imagenDecodificada image.Image

	imgExtension := filepath.Ext(rutaImagen)

	if imgExtension == ".jpg" || imgExtension == ".jpeg" {

		imagenParaDecodificar, imagenParaDecodificarError := os.Open(rutaImagen)
		if imagenParaDecodificarError != nil {
			return 0, 0, imagenDecodificada, imagenParaDecodificarError
		}

		decodificar, decodificarError := jpeg.Decode(imagenParaDecodificar)
		if decodificarError != nil {
			return 0, 0, imagenDecodificada, decodificarError
		}
		imagenParaDecodificar.Close()

		imagenParaDecodificarConfig, imagenParaDecodificarConfigError := os.Open(rutaImagen)
		if imagenParaDecodificarConfigError != nil {
			return 0, 0, imagenDecodificada, imagenParaDecodificarConfigError
		}

		decodificarConfig, decodificarConfigError := jpeg.DecodeConfig(imagenParaDecodificarConfig)
		if decodificarConfigError != nil {
			return 0, 0, imagenDecodificada, decodificarConfigError
		}
		imagenParaDecodificarConfig.Close()

		return decodificarConfig.Width, decodificarConfig.Height, decodificar, nil

	}
	if imgExtension == ".png" {

		imagenParaDecodificar, imagenParaDecodificarError := os.Open(rutaImagen)
		if imagenParaDecodificarError != nil {
			return 0, 0, imagenDecodificada, imagenParaDecodificarError
		}

		decodificar, decodificarError := jpeg.Decode(imagenParaDecodificar)
		if decodificarError != nil {
			return 0, 0, imagenDecodificada, decodificarError
		}
		imagenParaDecodificar.Close()

		imagenParaDecodificarConfig, imagenParaDecodificarConfigError := os.Open(rutaImagen)
		if imagenParaDecodificarConfigError != nil {
			return 0, 0, imagenDecodificada, imagenParaDecodificarConfigError
		}

		decodificarConfig, decodificarConfigError := png.DecodeConfig(imagenParaDecodificarConfig)
		if decodificarConfigError != nil {
			return 0, 0, imagenDecodificada, decodificarConfigError
		}
		imagenParaDecodificarConfig.Close()

		return decodificarConfig.Width, decodificarConfig.Width, decodificar, nil

	}

	return 0, 0, imagenDecodificada, errors.New("extensión no valida")

}

func IniciarProcesamiento(rutaImagenMuestra string, imagenesAnalizar []string) RespuestaProcesamiento {
	var respuestaProcesamiento RespuestaProcesamiento

	var inputRekognition rekognition.CompareFacesInput

	conexionRekognition := conectarRekognition()

	infoMuestra, infoMuestraError := convertirRekognitionImage(rutaImagenMuestra)
	if infoMuestraError != nil {
		respuestaProcesamiento.ErrorCritico = true
		respuestaProcesamiento.ErrorCriticoMensaje = infoMuestraError.Error()
		return respuestaProcesamiento
	}

	inputRekognition.SourceImage = &infoMuestra

	for _, imagen := range imagenesAnalizar {
		var resultadoProcesamiento ResultadoProcesamiento
		resultadoProcesamiento.Ruta = imagen
		resultadoProcesamiento.RutaDestino = filepath.Dir(rutaImagenMuestra)
		infoImagen, infoImagenError := convertirRekognitionImage(imagen)
		if infoImagenError != nil {
			resultadoProcesamiento.Error = true
			resultadoProcesamiento.ErrorMensaje = infoImagenError.Error()
			respuestaProcesamiento.Resultado = append(respuestaProcesamiento.Resultado, resultadoProcesamiento)

			continue
		}

		inputRekognition.TargetImage = &infoImagen

		comparacion, comparacionError := conexionRekognition.CompareFaces(&inputRekognition)

		if comparacionError != nil {
			resultadoProcesamiento.Error = true
			resultadoProcesamiento.ErrorMensaje = comparacionError.Error()
			respuestaProcesamiento.Resultado = append(respuestaProcesamiento.Resultado, resultadoProcesamiento)

			continue
		}

		if len(comparacion.FaceMatches) == 0 {
			respuestaProcesamiento.Resultado = append(respuestaProcesamiento.Resultado, resultadoProcesamiento)

			continue
		}

		resultadoProcesamiento.CarasDetectadas = len(comparacion.FaceMatches)

		for _, v := range comparacion.FaceMatches {
			var cuadroDelimitador CuadroDelimitador
			cuadroDelimitador.Height = *v.Face.BoundingBox.Height
			cuadroDelimitador.Left = *v.Face.BoundingBox.Left
			cuadroDelimitador.Top = *v.Face.BoundingBox.Top
			cuadroDelimitador.Width = *v.Face.BoundingBox.Width
			resultadoProcesamiento.CuadrosDelimitadores = append(resultadoProcesamiento.CuadrosDelimitadores, cuadroDelimitador)
		}

		ancho, alto, imagenDecodificada, obtenerInformacionError := obtenerInformacionImagen(imagen)
		if obtenerInformacionError != nil {
			resultadoProcesamiento.Error = true
			resultadoProcesamiento.ErrorMensaje = fmt.Sprintf("no se han conseguido la información (ancho, alto, decodificacion) del archivo: %v", obtenerInformacionError.Error())
			respuestaProcesamiento.Resultado = append(respuestaProcesamiento.Resultado, resultadoProcesamiento)

			continue
		}

		resultadoProcesamiento.ImagenDecodificada = imagenDecodificada

		resultadoProcesamiento.Ancho = ancho
		resultadoProcesamiento.Alto = alto

		resultadoProcesamiento.cuadrosDelimitadoresACoordenadas()
		errorPintar := resultadoProcesamiento.pintarDetecciones()
		if errorPintar != nil {
			resultadoProcesamiento.Error = true
			resultadoProcesamiento.ErrorMensaje = errorPintar.Error()
		}

		respuestaProcesamiento.Resultado = append(respuestaProcesamiento.Resultado, resultadoProcesamiento)

	}

	return respuestaProcesamiento

}
