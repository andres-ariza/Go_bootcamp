package main

import "fmt"

func main() {
	result := promedio(5, 2, 3)
	if result == -1 {
		fmt.Println("Error: ¡Ingresa un numero positivo!")
	} else {
		fmt.Println("El promedio de notas es: ", result)
	}
}

func promedio(notas ...int) float32 {
	var resultado float32
	for _, nota := range notas {
		resultado += float32(nota)
		if nota < 0 {
			return -1
		}
	}
	return resultado / float32(len(notas))
}
