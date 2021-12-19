package main

import (
	"errors"
	"fmt"
)

func main() {
	const (
		perro     = "perro"
		gato      = "gato"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)
	animalHamster, err := Animal(hamster)
	animalTarantula, err := Animal(tarantula)

	if err != nil {
		fmt.Println(err)
	} else {
		var cantidad float64
		cantidad += animalPerro(5)
		cantidad += animalGato(8)
		cantidad += animalHamster(4)
		cantidad += animalTarantula(3)

		fmt.Println("La cantidad de alimento necesario es: ", cantidad, "Kg")
	}

}

func Animal(tipo string) (func(cantidad float64) float64, error) {
	switch tipo {
	case "perro":
		return comidaPerro, nil
	case "gato":
		return comidaGato, nil
	case "hamster":
		return comidaHamster, nil
	case "tarantula":
		return comidaTarantula, nil
	default:
		return nil, errors.New("Animal incorrecto")
	}
}

func comidaPerro(cantidad float64) float64 {
	return cantidad * 10
}
func comidaGato(cantidad float64) float64 {
	return cantidad * 5
}
func comidaHamster(cantidad float64) float64 {
	return cantidad * 0.25
}
func comidaTarantula(cantidad float64) float64 {
	return cantidad * 0.15
}
