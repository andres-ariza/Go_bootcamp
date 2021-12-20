package main

import (
	"io/ioutil"
	"log"
)

func main() {
	producto := "id: 1, nombre: cepillo, precio: 3$, cantidad: 20;\nid: 2, nombre: shampoo, precio: 5$, cantidad: 10;\nid: 3, nombre: jabon, precio: 4$, cantidad: 25;\n"
	contenido := []byte(producto)
	err := ioutil.WriteFile("./listadoProductos.csv", contenido, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
