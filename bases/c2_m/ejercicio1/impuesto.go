package main

import "fmt"

func main() {
	fmt.Println("Deposito: ")
	var deposito float32
	fmt.Scan(&deposito)
	fmt.Print("Descuento por impuesto: $", impuesto(deposito), " \n")
}

func impuesto(deposito float32) float32 {
	if deposito > 50000 {
		return deposito * 0.17
	} else if deposito > 150000 {
		return deposito * 0.27
	} else {
		return 0
	}
}
