package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando...")
	data, err := os.ReadFile("customers.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(string(data))
	}
	fmt.Println("ejecución finalizada")
}
