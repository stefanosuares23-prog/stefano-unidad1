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

func bytesPorUnidad(unidad string) (float64, bool) {
	switch unidad {
	case "B":
		return 1, true
	case "KB":
		return 1024, true
	case "MB":
		return 1024 * 1024, true
	case "GB":
		return 1024 * 1024 * 1024, true
	default:
		return 0, false
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Uso: convertir <valor> <unidad_origen> <unidad_destino>")
		os.Exit(ErrUso)
	}

	valor, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil || math.IsNaN(valor) || math.IsInf(valor, 0) || valor < 0 {
		fmt.Fprintf(os.Stderr, "Valor invalido: %s\n", os.Args[1])
		os.Exit(ErrValor)
	}

	origen := strings.ToUpper(os.Args[2])
	bytesOrigen, okOrigen := bytesPorUnidad(origen)
	if !okOrigen {
		fmt.Fprintf(os.Stderr, "Unidad de origen invalida: %s\n", os.Args[2])
		os.Exit(ErrUnidad)
	}

	destino := strings.ToUpper(os.Args[3])
	bytesDestino, okDestino := bytesPorUnidad(destino)
	if !okDestino {
		fmt.Fprintf(os.Stderr, "Unidad de destino invalida: %s\n", os.Args[3])
		os.Exit(ErrUnidad)
	}

	resultado := (valor * bytesOrigen) / bytesDestino
	fmt.Printf("%g %s = %.2f %s\n", valor, origen, resultado, destino)
}
