package main

import (
	"fmt"
	"guia4/ejercicios"
	"guia4/set"
)

func main() {
	fmt.Println("Letras:")
	println(ejercicios.Letras("Hola mundo").String())
	println(ejercicios.Letras("Paralepipedo").String())
	println(ejercicios.Letras("Abracadabra").String())
	fmt.Println()

	s1 := set.NewSet(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	s2 := set.NewSet(5, 6, 7, 8, 10, 11, 12, 13, 1, 2, 3, 4, 17, 22)
	s3 := []int{1, 1, 2, 3, 4, 5, 5, 6, 6, 6, 6, 1, 1, 9, 8, 10, 10, 11, 10}

	fmt.Println("Interseccion: ", ejercicios.Interseccion(s1, s2).String())
	fmt.Println("Diferencia Simetrica: ", ejercicios.DiferenciaSimetrica(s1, s2).String())
	fmt.Println("Eliminar Repetidos: ", ejercicios.EliminarRepetidos(s3))

}
