package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Edad: ")
	var edad uint
	fmt.Scan(&edad)
	if edad < 22 {
		fmt.Println("No aplica al prestamo")
		os.Exit(0)
	}

	fmt.Println("¿Empleado?(y/n): ")
	var empleado string
	fmt.Scan(&empleado)
	if empleado != "y" {
		fmt.Println("No aplica al prestamo")
		os.Exit(0)
	}

	fmt.Println("Antiguedad(años): ")
	var antiguedad uint
	fmt.Scan(&antiguedad)
	if antiguedad <= 1 {
		fmt.Println("No aplica al prestamo")
		os.Exit(0)
	}

	fmt.Println("Sueldo: ")
	var sueldo uint
	fmt.Scan(&sueldo)
	if sueldo > 100000 {
		fmt.Println("Felicidades, puede aplicar al prestamo sin interes")
		os.Exit(0)
	}

	fmt.Println("Puede aplicar al prestamo")
	os.Exit(0)
}
