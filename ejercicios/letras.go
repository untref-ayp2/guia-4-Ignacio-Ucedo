package ejercicios

import (
	"guia4/set"
)

// recibe una cadena y devuelve el conjunto (set) de todas las letras de la cadena
// (aflabeto inglés, mayusculas y minúsculas). O(n^2)
func Letras(s string) *set.Set[string] {

	// var setPointer *set.Set[string]
	var set set.Set[string]

	for i := 0; i < len(s); i++ {

		if ((s)[i] >= 97 && (s)[i] <= 122) ||
			((s)[i] >= 65 && (s)[i] <= 90) { //ascii define 128 caracteres y las letras minúsculas
			//están entre los índices 97 y 122 inclusive. Y las mayusculas desde el 65 al 90. (indices en decimal)

			set.Add(string([]rune(s)[i])) //agrego letra por letra (char por char)
		}
	}

	return &set
}
