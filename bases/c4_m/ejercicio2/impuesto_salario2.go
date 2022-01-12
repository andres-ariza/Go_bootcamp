package main

import (
	"errors"
	"fmt"
)

func errorTest(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return errors.New("debe pagar impuesto")
	}
}

func main() {
	salary := 400000

	err := errorTest(salary)

	fmt.Println(err)
}
