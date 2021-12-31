package main

import "fmt"

func main() {
	fmt.Println("El salario es: $", salario(600, "C"))
}

func salario(minutos int, categoria string) float32 {
	switch categoria {
	case "C":
		return float32(minutos / 60 * 1000)
	case "B":
		return float32(minutos/60*1500) * 1.2
	case "A":
		return float32(minutos/60*3000) * 1.5
	}
	return -1
}
