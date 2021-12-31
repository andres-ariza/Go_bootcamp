package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func (m Matrix) Set() {
	if len(m.valores) != m.ancho*m.alto {
		fmt.Println("La cantidad de valores no coincide con las dimensiones especificadas")
	} else {
		fmt.Println("La cantidad de valores coincide con las dimensiones")

	}
}

func (m Matrix) Print() {
	if len(m.valores) == 0 {
		fmt.Println("La matriz esta vacia")
		return
	}
	for fila := 0; fila < m.alto; fila++ {
		fmt.Println(m.valores[fila*m.ancho : fila*m.ancho+m.ancho])
	}
}

func (m Matrix) isCuadratic() {
	if (m.alto == m.ancho) && m.alto != 0 {
		fmt.Println("Si, es cuadratica")
	} else {
		fmt.Println("No es cuadratica")

	}
}

func (m Matrix) Max() {
	max := -math.MaxFloat64
	for _, elemento := range m.valores {
		if elemento > max {
			max = elemento
		}
	}
	fmt.Println(max)
}
func main() {
	m := Matrix{
		valores: []float64{1, 2, 3, 4, 54, 65, 76, 87, 87},
		alto:    3,
		ancho:   3,
	}
	Matrix.Set(m)
	Matrix.Print(m)
	Matrix.isCuadratic(m)
	Matrix.Max(m)
}
