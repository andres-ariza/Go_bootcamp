package main

import (
	"fmt"
)

func errorTest(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de: %v", salary)
	} else {
		return fmt.Errorf("debe pagar impuesto")
	}
}

func main() {
	salary := 40000

	err := errorTest(salary)

	fmt.Println(err)
}
