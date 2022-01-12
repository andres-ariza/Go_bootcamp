package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26,
		"Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("\nLa edad de Benjamin es: ", employees["Benjamin"]) // Edad de Benjamin
	fmt.Println("-------\n", "Los empleados mayores de 21 años son: ")
	for key, element := range employees {
		if element > 21 { // Mayores a 21
			fmt.Println("- ", key)
		}
	}
	fmt.Println("-------\n", employees)
	employees["Federico"] = 25 // Agregar a federico
	fmt.Println("-------\n", employees)
	delete(employees, "Pedro") // Eliminar a pedro
	fmt.Println("-------\n", employees)
	fmt.Println()
}
