package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		minimo   = "minimo"
		promedio = "promedio"
		maximo   = "maximo"
	)
	minFunc, err1 := operacion(minimo)
	promFunc, err2 := operacion(promedio)
	maxFunc, err3 := operacion(maximo)

	if err1 != nil {
		fmt.Println(err1)
	} else if err2 != nil {
		fmt.Println(err2)
	} else if err3 != nil {
		fmt.Println(err3)
	} else {
		valorMinimo := minFunc(2, 3, 3, 4, 1, 2, 4, 5)
		valorPromedio := promFunc(2, 3, 3, 4, 5, 2, 4, 5)
		valorMaximo := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println("El valor minimo es: ", valorMinimo)
		fmt.Println("El valor promedio es: ", valorPromedio)
		fmt.Println("El valor maximo es: ", valorMaximo)
	}
}

func operacion(shape string) (func(notas ...float64) float64, error) {
	switch shape {
	case "minimo":
		return calculoMinimo, nil
	case "promedio":
		return calculoPromedio, nil
	case "maximo":
		return calculoMaximo, nil
	default:
		return nil, errors.New("Operacion invalida")
	}
}

func calculoMinimo(notas ...float64) float64 {
	var min float64
	for i, nota := range notas {
		if i == 0 {
			min = nota
		} else if min > nota {
			min = nota
		}
	}
	return min
}

func calculoMaximo(notas ...float64) float64 {

	var max float64
	for i, nota := range notas {
		if i == 0 {
			max = nota
		} else if max < nota {
			max = nota
		}
	}
	return max
}

func calculoPromedio(notas ...float64) float64 {
	var prom float64
	for _, nota := range notas {
		prom += nota
	}
	return prom / float64(len(notas))
}
