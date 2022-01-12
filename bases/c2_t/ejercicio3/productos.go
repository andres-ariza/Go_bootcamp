package main

import (
	"fmt"
)

// estructuras -------
type tiendaBog struct {
	nombre string
	precio float64
	tienda string
}
type tiendaMed struct {
	precio float64
	tienda string
}

// metodos de estructuras -------
func (t tiendaBog) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return t.precio
	case "mediano":
		var porcentaje float64
		porcentaje = (t.precio / 100) * 3
		return t.precio + porcentaje
	case "grande":
		var porcentaje float64
		flete := 2500
		porcentaje = (t.precio / 100) * 6
		return t.precio + porcentaje + float64(flete)
	}
	return t.precio
}

func (t tiendaBog) Envio(dir string) string {
	enviado := "Enviando un paquete a " + dir
	return enviado
}
func (t tiendaMed) Precio(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return t.precio
	case "mediano":
		var porcentaje float64
		porcentaje = (t.precio / 100) * 3
		return t.precio + porcentaje
	case "grande":
		var porcentaje float64
		flete := 2500
		porcentaje = (t.precio / 100) * 6
		return t.precio + porcentaje + float64(flete)
	}
	return t.precio
}
func (t tiendaMed) Envio(dir string) string {
	enviado := "Enviando un paquete a " + dir
	return enviado
}

// interfaces -------
type Ecommerce interface {
	Precio(size string) float64
	Envio(direccion string) string
}

// funciones de interfaces --------
func nuevaTienda(mailType string) Ecommerce {
	if mailType == "Tienda Rola" {
		return tiendaBog{nombre: "mesa", precio: 2000, tienda: "sucursal bogota"}
	}
	if mailType == "Tienda Paisa" {
		return tiendaMed{precio: 5000, tienda: "sucursal medellin"}
	}
	return nil
}
func main() {
	fmt.Println("------------------------------------------------")

	tiendaMedellin := nuevaTienda("Tienda Paisa")          // func interfase
	precioMedellin := tiendaMedellin.Precio("grande")      // struct method
	envioMedellin := tiendaMedellin.Envio("Parque lleras") // struct method

	fmt.Println(tiendaMedellin)
	fmt.Println(envioMedellin)
	fmt.Println("Precio del producto:", precioMedellin)

	fmt.Println("------------------------------------------------")

	tiendaBogota := nuevaTienda("Tienda Rola")     // func interfase
	precioBogota := tiendaBogota.Precio("Mediano") // struct method
	envioBogota := tiendaBogota.Envio("Calle 100") // struct method

	fmt.Println(tiendaBogota)
	fmt.Println(envioBogota)
	fmt.Println("Precio del pruducto:", precioBogota)

	fmt.Println("------------------------------------------------")
}
