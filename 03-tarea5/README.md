# 03-tarea5 — Analizador de texto y límite de subtítulos

## Descripción
Herramienta de línea de comandos escrita en Go que analiza un texto recibido por argumento y reporta su cantidad de bytes, runas, palabras, letras, dígitos, mayúsculas y runas multibyte (UTF-8), además de validar si entra dentro del límite de 42 runas para subtítulos.

## Cómo ejecutar
Estando en la carpeta `03-tarea5`:

```bash
go mod init tarea5
go build -o analizar.exe .
./analizar.exe "Canción de cuna 2026"