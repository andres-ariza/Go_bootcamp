package main

import (
	"fmt"
)

func main() {
	numero := 1
	mes := []string{"Enero", "Febrero", "Marzo", "Abril",
		"Mayo", "Junio", "Julio", "Agosto",
		"Septiembre", "Octubre", "Noviembre", "Diciembre"}
	fmt.Println(mes[numero-1])
}

// Tambien se podria resolver con case, elijo esta forma por que se repiten menos prints, o con map pero tendria que escribir todos los indices con un for o algo
