// Package almacenamiento provee constantes y funciones para convertir unidades de almacenamiento (B, KB, MB y GB).
package almacenamiento

import "math"

// B, KB, MB y GB definen el tamaño en bytes de cada unidad usando base 1024.
const (
    B  = 1.0
    KB = 1024 * B
    MB = 1024 * KB
    GB = 1024 * MB
)

// Bytes devuelve cuántos bytes vale la unidad indicada en mayúsculas y si dicha unidad es válida.
func Bytes(unidad string) (float64, bool) {
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

// Convertir convierte un valor desde bytesOrigen hacia bytesDestino y redondea el resultado a dos decimales.
func Convertir(valor, bytesOrigen, bytesDestino float64) float64 {
    return redondear(valor * bytesOrigen / bytesDestino)
}

func redondear(v float64) float64 {
    return math.Round(v*100) / 100
}
