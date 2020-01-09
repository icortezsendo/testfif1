package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSeparaDatos(t *testing.T) {
	lista := [9][2]string{{"A0511AB398765UJ1N230200", "23"},
		{"C0511AB398765UJ1N230200", "Tipo dato debe ser N o A, no C"},
		{"N0511AB398765UJ1N230200", "Tipo de dato definido como Numerico y es Alfanumerico"},
		{"A051cAB398765UJ1N230200", "Error en convertir el largo a numerico: 1c"},
		{"A0511AB398765UJ1C230200", "Tipo dato debe ser N o A, no C"},
		{"A0511AB398765UJ1N230c00", "Error en convertir el largo a numerico: 0c"},
		{"A0511AB398765UJ1N23020c", "Tipo de dato definido como Numerico y es Alfanumerico"},
		{"", "No se introdujo el dato a leer"},
		{"A0511AB398765UJ1N230200N66011", "23"},
	}
	for _, dato := range lista {
		mapa, err := separaDatos([]byte(dato[0]))
		if err == "" {
			assert.Contains(t, mapa, dato[1])
		} else {
			assert.Equal(t, err, dato[1])
		}

	}
}
