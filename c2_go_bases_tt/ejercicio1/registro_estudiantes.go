package main

import (
	"fmt"
)

func main() {
	type Fecha struct {
		dia int
		mes string
		año int
	}
	type Alumnos struct {
		Nombre   string
		Apellido string
		DNI      int
		Fecha    Fecha
	}

	p1 := Alumnos{"Pepe", "Contreras", 198278345, Fecha{04, "Enero", 1998}}
	detalle(p1)
}

func detalle(a Alumnos) {
	fmt.Println(p1)
}
