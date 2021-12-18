package main

import (
	"fmt"
)

func main() {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos",
		"Pedro", "Axel", "Alez", "Dolores",
		"Federico", "Hernán", "Leandro",
		"Eduardo", "Duvraschka"}
	estudiantes = append(estudiantes, "Gabriela")
	fmt.Println(estudiantes)
}
