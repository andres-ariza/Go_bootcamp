package main

import (
	"fmt"
)

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func cambiarNombre(n *string, a *string) {

}
func cambiarEdad(e *int) {

}
func cambiarCorreo(c *string) {

}
func cambiarContraseña(c *string) {

}

func main() {
	user := usuario{"Andres", "Ariza", 27, "afaarizap@gmail.com", "123abc"}
	var p *usuario

	nuevoNombre, nuevoApellido := "Alice", "Medrano"

	p = &user
	cambiarNombre(&p.Nombre, &p.Apellido)
	cambiarEdad(&p.Edad)
	cambiarCorreo(&p.Correo)
	cambiarContraseña(&p.Contraseña)

	fmt.Println(user.Apellido)
	print()

}
