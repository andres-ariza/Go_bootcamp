package main

import (
	"fmt"
)

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func (m Matrix) Set() {
	if len(m.valores) != m.ancho*m.alto {
		fmt.Println("La cantidad de valores no coincide con las dimensiones especificadas")
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
		fmt.Println("True")
	}
	fmt.Println("False")
}

func (m Matrix) Max() {

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
