package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	// preparacion
	numerador := 5
	denominador := 0

	//accion
	_, err := Dividir(numerador, denominador)

	//verificacion
	assert.NotNil(t, err)
}
func TestDividir2(t *testing.T) {
	// preparacion
	numerador := 9
	denominador := 3
	resultadoEsperado := 3

	//accion
	resultado, _ := Dividir(numerador, denominador)

	//verificacion
	assert.Equal(t, resultado, resultadoEsperado, "deberian ser iguales")
}
