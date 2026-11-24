package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

const (
    ErrUso    = 1
    ErrValor  = 2
    ErrUnidad = 3
)

const (
    B  = 1.0
    KB = 1024 * B
    MB = 1024 * KB
    GB = 1024 * MB
)

func bytesUnidad(unidad string) (float64, bool) {
    switch unidad {
    case "B":
        return B, true
    case "KB":
        return KB, true
    case "MB":
        return MB, true
    case "GB":
        
        return GB, true
    default:
        return 0, false
    }
}

func redondear2(v float64) float64 {
    return math.Round(v*100) / 100
}

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

    bytesOrigen, okOrigen := bytesUnidad(origen)
    bytesDestino, okDestino := bytesUnidad(destino)
    if !okOrigen || !okDestino {
        fmt.Fprintln(os.Stderr, "Error: unidad inválida (use B, KB, MB o GB)")
        os.Exit(ErrUnidad)
    }

    resultado := redondear2(valor * bytesOrigen / bytesDestino)
    fmt.Printf("%s %s = %.2f %s\n", valorStr, origen, resultado, destino)
}
