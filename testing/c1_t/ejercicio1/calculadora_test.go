package calculadora

import (
	"testing"
)

func TestRestar(t *testing.T) {
	// preparacion
	a := 5
	b := 3
	resultadoEsperado := 2

	//accion
	resultado := Restar(a, b)

	//verificacion
	if resultado != resultadoEsperado {
		t.Errorf("Funcion resta() devolvio: %v, pero el valor esperado es: %v", resultado, resultadoEsperado)
	}
}
