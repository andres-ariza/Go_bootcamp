package main

import "fmt"

// structs
type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}
type Producto struct {
	Nombre   string
	Precio   float32
	Cantidad int
}

// funciones
func nuevoProducto(Name string, Price float32) Producto {
	return Producto{
		Nombre:   Name,
		Precio:   Price,
		Cantidad: 0,
	}
}
func agregarProducto(User *Usuario, Product *Producto, Quantity int) {
	User.Productos = append(User.Productos, *Product)

}
func borrarProductos(User *Usuario) {
	User.Productos = nil
}
func main() {
	producto := nuevoProducto("leche", 3200.0)
	fmt.Println("leche: ", producto)
	cliente := Usuario{
		Nombre:    "Andres",
		Apellido:  "Ariza",
		Correo:    "afarizap@gmail.com",
		Productos: []Producto{},
	}
	fmt.Println("cliente: ", cliente)
	fmt.Printf("productos: %v %T\n", cliente.Productos, cliente.Productos)
	agregarProducto(&cliente, &producto, 3)
	fmt.Println("cliente+producto: ", cliente)
	borrarProductos(&cliente)
	fmt.Println("cliente-productos", cliente)
}
