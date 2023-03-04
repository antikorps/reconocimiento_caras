package procesar

import "image"

type RespuestaProcesamiento struct {
	Resultado           []ResultadoProcesamiento `json:"resultado"`
	ErrorCritico        bool                     `json:"errorCritico"`
	ErrorCriticoMensaje string                   `json:"errorCriticoMensaje"`
}

type CuadroDelimitador struct {
	Height float64
	Left   float64
	Top    float64
	Width  float64
}

type ResultadoProcesamiento struct {
	Ancho                  int                 `json:"-"`
	Alto                   int                 `json:"-"`
	CuadrosDelimitadores   []CuadroDelimitador `json:"-"`
	CoordenadasPintar      []image.Point       `json:"-"`
	CarasDetectadas        int                 `json:"carasDetectadas"`
	Error                  bool                `json:"error"`
	ErrorMensaje           string              `json:"errorMensaje"`
	ImagenDecodificada     image.Image         `json:"-"`
	Ruta                   string              `json:"ruta"`
	RutaDestino            string              `json:"-"`
	ImagenResultanteNombre string              `json:"nombre"`
	CoordenadasManuales    [][]int64           `json:"-"`
}
