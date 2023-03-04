export type {imagenesExploradas, analisisExploracion, respuestaProcesamiento, resultadoProcesamiento}

interface imagenesExploradas {
    validas: analisisExploracion[],
    invalidas: analisisExploracion[]
}

interface analisisExploracion {
    ruta: string,
    error: string,
}

interface respuestaProcesamiento {
    resultado: resultadoProcesamiento[],
    errorCritico: boolean,
    errorCriticoMensaje: string
}

interface resultadoProcesamiento {
    carasDetectadas: number,
    error: boolean,
    errorMensaje: string,
    ruta: string,
    nombre: string
}