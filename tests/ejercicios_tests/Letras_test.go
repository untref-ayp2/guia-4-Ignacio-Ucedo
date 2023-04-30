package tests

import (
	"guia4/ejercicios"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetras(t *testing.T) {
	letras := ejercicios.Letras("Hola Mundo")
	if letras.Size() != 8 {
		t.Errorf("Letras(\"Hola Mundo\") debería tener 8 caracteres")
	}

	assert.Equal(t, 8, letras.Size())

	if !letras.Contains("H") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra H")
	}
	if !letras.Contains("o") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra o")
	}
	if !letras.Contains("l") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra l")
	}
	if !letras.Contains("a") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra a")
	}
	if !letras.Contains("M") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra M")
	}
	if !letras.Contains("u") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra u")
	}
	if !letras.Contains("n") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra n")
	}
	if !letras.Contains("d") {
		t.Errorf("Letras(\"Hola Mundo\") debería contener la letra d")
	}

}
