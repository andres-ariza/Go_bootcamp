package main

import (
	"fmt"
)

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func cambiarNombre(Name *string, lName *string) {
	*Name = "Alice"
	*lName = "Medrano"
}
func cambiarEdad(Age *int) {
	*Age = 26
}
func cambiarCorreo(Mail *string) {
	*Mail = "alice.medrani@gmail.com"
}
func cambiarContraseña(Pw *string) {
	*Pw = "asd123"
}

func main() {
	p := Usuario{"Andres", "Ariza", 27, "afarizap@gmail.com", "123abc"}
	fmt.Println("Los valores del usuario son: ", p)

	cambiarNombre(&p.Nombre, &p.Apellido)
	cambiarEdad(&p.Edad)
	cambiarCorreo(&p.Correo)
	cambiarContraseña(&p.Contraseña)

	fmt.Println("Los valores del usuario ahora son: ", p)
	print()

}
