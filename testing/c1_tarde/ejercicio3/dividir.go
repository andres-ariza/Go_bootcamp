package dividir

import "errors"

func Dividir(numerador, denominador int) (int, error) {
	if denominador == 0 {
		return 0, errors.New("el denominador no puede ser igual a 0")
	}
	return numerador / denominador, nil
}
