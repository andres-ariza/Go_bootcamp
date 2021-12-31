package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    Fecha
}
type Fecha struct {
	dia int
	mes string
	año int
}

func (a Alumno) detalle() {
	fmt.Printf("Detalle del alumno: %+v\n", a)

}
func main() {

	p1 := Alumno{"Pepe", "Contreras", 198278345, Fecha{04, "Enero", 1998}}
	Alumno.detalle(p1)
}
