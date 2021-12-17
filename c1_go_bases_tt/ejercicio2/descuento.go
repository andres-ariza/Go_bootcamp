package main

import "fmt"

func main() {
	fmt.Println("Ingresa precio del producto: ")
	var producto float32
	fmt.Scan(&producto)

	fmt.Println("Ingresa descuento: ")
	var descuento float32
	fmt.Scan(&descuento)

	fmt.Println("Precio con descuento es: ", producto*(1-(descuento/100)))
}
