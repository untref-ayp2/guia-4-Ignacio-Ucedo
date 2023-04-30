package ejercicios

import (
	"guia4/set"
)

// devuelve un conjunto con los elementos de s1 y s2 que sólo pertenecen a alguno de ellos.
// (O(n^2))
func DiferenciaSimetrica[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {

	//para no tocar los set de entrada:
	a := set.NewSet(s1.Values()...)
	b := set.NewSet(s2.Values()...)

	for i := 0; i < b.Size(); i++ {
		if a.Contains(b.Values()[i]) {
			a.Remove(b.Values()[i])
		} else {
			a.Add(b.Values()[i])
		}

	}

	return a
}
