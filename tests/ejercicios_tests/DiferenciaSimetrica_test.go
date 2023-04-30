package tests

import (
	"guia4/ejercicios"
	"guia4/set"
	"testing"
)

func TestDiferenciaSimetrica(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(3, 4, 5, 6)
	s3 := ejercicios.DiferenciaSimetrica(s1, s2)
	if s3.Size() != 4 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 4)
	}
}

func TestNilpotencia(t *testing.T) {
	s1 := set.NewSet(1, 2, 3, 4)
	s2 := set.NewSet(1, 2, 3, 4)
	s3 := ejercicios.DiferenciaSimetrica(s1, s2)
	if s3.Size() != 0 {
		t.Errorf("Tamaño %d, debería ser: %d", s3.Size(), 0)
	}
}

func TestAsociativa(t *testing.T) {
	a := set.NewSet(1, 2, 3, 4)
	b := set.NewSet(1, 2)
	c := set.NewSet(6, 7, 2)

	difAyB := ejercicios.DiferenciaSimetrica(a, b) // 3 , 4
	difByC := ejercicios.DiferenciaSimetrica(b, c) // 1 , 6 , 7

	e1 := ejercicios.DiferenciaSimetrica(difAyB, c) // 3 , 4 , 6 , 7 , 2
	e2 := ejercicios.DiferenciaSimetrica(difByC, a) // 2 , 3 , 4, 6 , 7

	if e1.Size() != e2.Size() {
		t.Errorf("Tamaño de e1 = %d, y de e2 = %d", e1.Size(), e2.Size())
	}
}

func TestConmutativa(t *testing.T) {
	a := set.NewSet(1, 2, 3, 4)
	b := set.NewSet(1, 2)

	e1 := ejercicios.DiferenciaSimetrica(a, b)
	e2 := ejercicios.DiferenciaSimetrica(b, a)

	if e1.Size() != e2.Size() {
		t.Errorf("Tamaño de e1 = %d, y de e2 = %d", e1.Size(), e2.Size())
	}
}

func TestNulo(t *testing.T) {

	a := set.NewSet(1, 2, 3, 4)
	b := set.NewSet(1)

	b.Remove(1)

	e := ejercicios.DiferenciaSimetrica(a, b)

	if e.Size() != a.Size() {
		t.Errorf("Tamaño de e = %d, y de a = %d", e.Size(), a.Size())
	}
}
