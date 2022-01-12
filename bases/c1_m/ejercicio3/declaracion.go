package main

import "fmt"

func main() {
	var nombre string   //incorrecta
	var apellido string // correcta
	var edad int        //incorrecta
	// 1apellido := 6 // repetida
	var licencia_de_conducir bool  //incorrecta
	var estatura_de_la_persona int //incorrecta
	var cantidadDeHijos int        //incorrecta

	fmt.Println("Solo la variable apellido es correcta")
	fmt.Println(nombre, apellido, edad, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)
}
