// Escribir una función que reciba una cantidad variables de conjuntos de elementos comparables y devuelva la intersección de todos ellos.
package ejercicios

import (
	"guia4/set"
)

func Interseccion[T comparable](s1, s2 *set.Set[T]) *set.Set[T] {
	s3 := set.NewSet[T]()
	a1 := s1.Values()
	a2 := s2.Values()

	for i := 0; i < s1.Size(); i++ {
		if s1.Contains(a1[i]) && s2.Contains(a1[i]) {
			s3.Add(a1[i])
		}
	}
	for i := 0; i < s2.Size(); i++ {
		if s1.Contains(a2[i]) && s2.Contains(a2[i]) {
			s3.Add(a2[i])
		}
	}

	return s3
}
