// Escribir una función que reciba una cadena y devuelva el conjunto de todas las letras de la cadena
package ejercicios

import (
	"guia4/set"
)

func Letras(s string) *set.Set[string] {

	set := set.NewSet[string]()

	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			set.Add(string(s[i]))
		}
	}
	return set
}
