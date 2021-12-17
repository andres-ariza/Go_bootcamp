package main

import "fmt"

func main() {
	var apellido string = "Gomez" //correcta
	var edad int = 35             // incorrecta
	var licencia bool             //incorrecta
	var sueldo float32 = 45857.90 //incorrecta
	var nombre string = "Julián"  //correcta

	fmt.Println(apellido, edad, licencia, sueldo, nombre)

}
