package main

import (
	"fmt"
)

type mensajeError struct {
	msg string
}

func (e mensajeError) Error() string {
	return e.msg
}

func errorTest(salary int) error {
	if salary < 150000 {
		return &mensajeError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
	} else {
		return mensajeError{
			msg: "Debe pagar impuesto",
		}
	}
}

func main() {
	salary := 400000

	err := errorTest(salary)

	fmt.Println(err)
}
