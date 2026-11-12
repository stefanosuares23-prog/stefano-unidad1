package main

import (
	"fmt"
)

type Permiso uint8

const (
	Leer Permiso = 1 << iota
	Escribir
	Ejecutar
)

func describir(p Permiso) string {
	s := ""
	if p&Leer != 0 {
		s += "r"
	} else {
		s += "-"
	}
	if p&Escribir != 0 {
		s += "w"
	} else {
		s += "-"
	}
	if p&Ejecutar != 0 {
		s += "x"
	} else {
		s += "-"
	}
	return s
}

func main() {
	usuario := Leer | Escribir | Ejecutar
	grupo := Leer | Ejecutar
	otros := Permiso(0)

	fmt.Printf("%s-%s-%s\n", describir(usuario), describir(grupo), describir(otros))
	fmt.Printf("numérico: %o%o%o\n", usuario, grupo, otros)

	otros |= Escribir
	grupo &^= Ejecutar

	fmt.Printf("%s-%s-%s\n", describir(usuario), describir(grupo), describir(otros))
}