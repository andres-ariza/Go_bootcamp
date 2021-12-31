package main

type tienda struct {
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	CalcularCosto()
}

type Ecomerce interface {
	Total()
	Agregar()
}

func (p producto) CalcularCosto(tamaño string) float64 {
	switch tamaño {
	case "pequeño":
		return p.precio
	case "mediano":
		var porcentaje float64
		porcentaje = p.precio * 0.03
		return p.precio + porcentaje
	case "grande":
		var porcentaje float64
		adicional := float64(2500)
		porcentaje = p.precio * 0.06
		return p.precio + porcentaje + adicional
	}
	return p.precio
}

func nuevaTienda(tipo string) Ecomerce {
	if tipo == "ropa" {
		return tienda{nombre: "pantalon", precio: 20000, tienda: "H&M", url: "http//H&M.com"}
	}
}

func nuevoProducto(tipo string, nombre string, precio float64) Producto {

}

func main() {
	justoybueno := nuevaTienda("ropa")
}
