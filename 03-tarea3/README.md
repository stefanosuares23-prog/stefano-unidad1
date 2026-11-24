# 03-tarea3 - Conversor de unidades de almacenamiento

## Descripcion
Herramienta CLI en Go que convierte valores entre unidades de almacenamiento (B, KB, MB y GB) validando argumentos, valores numericos no negativos y unidades permitidas con codigos de salida diferenciados (1, 2 y 3).

## Como ejecutar
Desde la carpeta 03-tarea3, compila con `go build -o convertir.exe .` y ejecuta por ejemplo `.\convertir.exe 1024 B KB`.

## Pruebas realizadas
- Conversion valida (`1024 B KB` -> `1.00 KB`, codigo de salida `0`).
- Cantidad incorrecta de argumentos (`10 KB` -> codigo de salida `1`).
- Valor invalido o negativo (`abc KB MB` o `-5 KB MB` -> codigo de salida `2`).
- Unidad de origen o destino invalida (`10 TB MB` o `10 KB TB` -> codigo de salida `3`).
- Formato y analisis estatico con `gofmt -l .` y `go vet ./...`.

## Preguntas
- Por que se separan los codigos de salida (`1`, `2`, `3`) y se envian los mensajes de error a `os.Stderr`? Para permitir que scripts o shells distingan entre errores de uso, datos invalidos o unidades no soportadas sin contaminar la salida estandar (`os.Stdout`).
