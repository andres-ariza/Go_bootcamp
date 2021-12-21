package main

import (
	"errors"
	"fmt"
)

func salary(horas, valor_hora int) (float64, error) {
	var impuesto float64
	salario := float64(horas * valor_hora)
	if salario >= 150000 {
		impuesto = salario * 0.1
		return impuesto, nil
	} else if salario < 80 {
		return impuesto, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		return 0, nil
	}
}

func bonus(salario float64, meses float64) (float64, error) {
	bono := salario / 12
	if meses != 0 {
		return bono * meses, nil
	} else {
		return 0, fmt.Errorf("el trabajador tiene %v meses", meses)
	}
}

func main() {
	horas := 960
	valor_hora := 4000
	meses := float64(1)

	salario, _ := salary(horas, valor_hora)
	aguinaldo, err := bonus(salario, meses)

	if err != nil {

	}

	fmt.Println("El salario del trabajador es: ", salario)
	fmt.Println("Sus aguinaldos son: ", aguinaldo, err)
}
