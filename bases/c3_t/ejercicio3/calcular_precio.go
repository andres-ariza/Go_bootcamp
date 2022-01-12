package main

import (
	"fmt"
)

// structs
type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}
type Servicio struct {
	Nombre  string
	Precio  float32
	Minutos int
}
type Mantenimiento struct {
	Nombre string
	Precio float32
}

// funciones
func sumarProductos(productos []Producto, c chan float32) {
	var result float32
	for _, v := range productos {
		result += float32(v.Cantidad) * v.Precio
	}
	c <- result
}

func sumarServicios(servicios []Servicio, c chan float32) {
	var result float32
	for _, v := range servicios {
		result += float32(v.Minutos) * v.Precio
	}
	c <- result
}
func sumarMantenimiento(mantenimientos []Mantenimiento, c chan float32) {
	var result float32
	for _, v := range mantenimientos {
		result += v.Precio
	}
	c <- result
}
func main() {
	var total float32
	// inicializar productos
	productos := []Producto{
		{
			Nombre:   "producto1",
			Precio:   20000.00,
			Cantidad: 3,
		},
		{
			Nombre:   "producto2",
			Precio:   60000.00,
			Cantidad: 2,
		},
	}
	c := make(chan float32)
	fmt.Println("los productos son: ", productos)
	go sumarProductos(productos, c)
	total += <-c
	fmt.Println("la suma va en: ", total)

	// inicializar servicios
	servicios := []Servicio{
		{
			Nombre:  "servicio1",
			Precio:  30000.00,
			Minutos: 3,
		},
		{
			Nombre:  "servicio2",
			Precio:  80000.00,
			Minutos: 2,
		},
	}
	fmt.Println("los servicios son: ", servicios)

	go sumarServicios(servicios, c)
	total += <-c
	fmt.Println("la suma va en: ", total)

	// inicializar mantenimientos
	mantenimientos := []Mantenimiento{
		{
			Nombre: "mantenimiento1",
			Precio: 10000.00,
		},
		{
			Nombre: "mantenimiento2",
			Precio: 30000.00,
		},
	}
	fmt.Println("los mantenimientos son: ", mantenimientos)

	go sumarMantenimiento(mantenimientos, c)
	total += <-c
	fmt.Println("la suma va en: ", total)

}
