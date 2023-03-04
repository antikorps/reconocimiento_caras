# Reconocimiento de caras

## Descripción

Programa para el reconocimiento de caras usando el servicio Rekognition de AWS

## Descargas
[GNU/Linux (amd64)](https://github.com/antikorps/reconocimiento_caras/raw/main/bin/linux/reconocimiento_caras-amd64-linux-amd64)
[Windows (amd64)](https://github.com/antikorps/reconocimiento_caras/raw/main/bin/windows/reconocimiento_caras-amd64.exe)

## Funcionamiento

El programa se distribuye compilado en un archivo autoejecutable para GNU/Linux y Windows (amd 64), por lo que no es necesario instalar nada. Para configurar el acceso a rekognition es necesario pasar las siguientes variables de entorno al programa:
```
AWS_ACCESS_KEY_ID=XXX
AWS_SECRET_ACCESS_KEY=XXX
AWS_SESSION_TOKEN=XXX
REGION=XXX
```
Aunque se pueden pasar manualmente, la forma más cómoda es crear un archivo .env en el mismo directorio que el ejectuable con esas opciones.

Cuando está en funcionamiento creará una carpeta llamada "imagenes_reconocimiento" en el mismo directorio que el ejecutable. En este directorio se encuentran las imágenes con las caras destacadas por si fueran necesarias.

![Funcionamiento](https://gifyu.com/images/evil_dead.gif)