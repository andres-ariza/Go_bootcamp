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
	} else if horas < 80 {
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
		return 0, fmt.Errorf("error, el trabajador tiene %v meses", meses)
	}
}

func main() {
	horas := 70
	valor_hora := 1
	meses := float64(0)

	salario, err1 := salary(horas, valor_hora)
	aguinaldo, err2 := bonus(salario, meses)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("El salario del trabajador es: ", salario)
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Sus aguinaldos son: ", aguinaldo)
	}
}
