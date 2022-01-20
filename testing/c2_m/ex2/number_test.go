package TypeNumber

import "testing"

type Test struct {
	in  int
	out string
}

var tests = []Test{{-1, "negative"}, {5, "positive"}}

func TestTypeNumber(t *testing.T) {
	for i, test := range tests {
		size := TypeNumber(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}

// go test -cover ./...
// muestra el coverage de nuestros test

// go test -cover -coverprofile=coverage.out  ./...
// genera un reporte en un archivo coveerage.out

// go tool cover -html=coverage.out
// genera una pestaña en explorador con las secciones testeadas resaltadas y las que no
