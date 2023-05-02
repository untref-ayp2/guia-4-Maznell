//Escribir una función que reciba un arreglo de elementos comparables y elimine los repetidos.

package ejercicios

import "guia4/set"

func EliminarRepetidos[T comparable](arreglo []T) []T {
	set := set.NewSet[T]()
	for i := 0; i < len(arreglo); i++ {
		if !set.Contains(arreglo[i]) {
			set.Add(arreglo[i])
		}
	}
	array := set.Values()
	return array
}
