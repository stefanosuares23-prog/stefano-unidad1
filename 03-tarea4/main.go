package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"

    "tarea4/internal/almacenamiento"
)

const (
    ErrGeneral = 1
    ErrUso     = 2
    ErrValor   = 3
    ErrUnidad  = 4
)

func main() {
    if len(os.Args) != 4 {
        fmt.Fprintln(os.Stderr, "Uso: go run . <valor> <unidad_origen> <unidad_destino>")
        os.Exit(ErrUso)
    }

    valorStr := os.Args[1]
    valor, err := strconv.ParseFloat(valorStr, 64)
    if err != nil || math.IsNaN(valor) || math.IsInf(valor, 0) || valor < 0 {
        fmt.Fprintln(os.Stderr, "Error: el valor debe ser un número mayor o igual a 0")
        os.Exit(ErrValor)
    }

    origen := strings.ToUpper(os.Args[2])
    destino := strings.ToUpper(os.Args[3])

    bytesOrigen, okOrigen := almacenamiento.Bytes(origen)
    bytesDestino, okDestino := almacenamiento.Bytes(destino)
    if !okOrigen || !okDestino {
        fmt.Fprintln(os.Stderr, "Error: unidad inválida (use B, KB, MB o GB)")
        os.Exit(ErrUnidad)
    }

    resultado := almacenamiento.Convertir(valor, bytesOrigen, bytesDestino)
    fmt.Printf("%s %s = %.2f %s\n", valorStr, origen, resultado, destino)
}
