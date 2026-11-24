package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	texto := "Programación en Go: rápido, simple y ¡divertido! Año 2026."

	numBytes := len(texto)
	numRunas := utf8.RuneCountInString(texto)

	palabras := strings.Fields(texto)
	numPalabras := len(palabras)

	mayusculas := 0
	digitos := 0
	for _, r := range texto {
		if unicode.IsUpper(r) {
			mayusculas++
		}
		if unicode.IsDigit(r) {
			digitos++
		}
	}

	letrasEspeciales := "áéíóúÁÉÍÓÚñÑ"
	conTildeON := 0
	for _, r := range texto {
		if strings.ContainsRune(letrasEspeciales, r) {
			conTildeON++
		}
	}

	fmt.Printf("Bytes: %d\n", numBytes)
	fmt.Printf("Runas: %d\n", numRunas)
	fmt.Printf("Palabras: %d\n", numPalabras)
	fmt.Printf("Mayúsculas: %d\n", mayusculas)
	fmt.Printf("Dígitos: %d\n", digitos)
	fmt.Printf("Con tilde/ñ: %d\n", conTildeON)

	fmt.Println("\n--- Explicación: ¿Por qué bytes y runas no coinciden? ---")
	fmt.Println("En Go, un 'string' es una secuencia de bytes de solo lectura (UTF-8).")
	fmt.Println("- Un byte equivale a 8 bits (tipo uint8). Los caracteres ASCII estándar ocupan 1 byte.")
	fmt.Println("- Una runa (tipo rune / int32) representa un punto de código Unicode. Las tildes y símbolos ocupan múltiples bytes.")
	fmt.Println("Por esta razón, la longitud en bytes (len) es mayor que el conteo real de caracteres (runas).")
}