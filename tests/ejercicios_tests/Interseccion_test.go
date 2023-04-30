package tests

import (
	"guia4/ejercicios"
	"guia4/set"
	"testing"
)

func TestInterseccion(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := ejercicios.Interseccion(s1, s2)
	if s3.Size() != 2 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 2)
	}
}
