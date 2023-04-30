package ejercicios

import (
	"guia4/set"
)

// Disminuir orden? (O(m*n^3))
func Interseccion[T comparable](conjuntos ...*set.Set[T]) *set.Set[T] {

	menor := conjuntos[0]

	for i := 0; i < len(conjuntos); i++ {
		if conjuntos[i].Size() < menor.Size() {
			menor = conjuntos[i]
		}
	}

	for i := 0; i < menor.Size(); i++ {
		for j := 0; j < len(conjuntos); j++ {
			if !conjuntos[j].Contains(menor.Values()[i]) {
				menor.Remove(menor.Values()[i])
				i--
			}
		}
	}

	return menor
}
