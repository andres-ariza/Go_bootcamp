package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProbar(t *testing.T) {
	// preparacion
	array := []int{5, 3, 1, 4, 2}
	resultadoEsperado := []int{1, 2, 3, 4, 5}

	//accion
	resultado := Probar(array)

	//verificacion
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
