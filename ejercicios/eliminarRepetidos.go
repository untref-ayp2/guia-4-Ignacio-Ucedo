package ejercicios

import "guia4/set"

//O(n^2)
func EliminarRepetidos[T comparable](arreglo []T) []T {

	var set set.Set[T]

	for i := 0; i < len(arreglo); i++ {
		set.Add(arreglo[i])
	}

	arreglo = make([]T, set.Size()) //ahora arreglo apunta a otro con valores nil y del largo del set.

	for i := 0; i < set.Size(); i++ {
		arreglo[i] = set.Values()[i]
	}

	return arreglo
}
