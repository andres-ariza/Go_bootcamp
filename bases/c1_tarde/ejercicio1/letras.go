package main

import "fmt"

func main() {
	var rae string = "Paralelepipedo"
	fmt.Println("Cantidad de letras:", len(rae))
	for i := 0; i < len(rae); i++ {
		fmt.Println(string(rae[i]))
	}
}
