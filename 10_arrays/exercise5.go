package main

import (
	"fmt"
	"math/rand"
)

func buscar(slice []int, buscado int) int {
	for i := range slice {
		if slice[i] == buscado {
			return i
		}
	}
	return -1
}

func buscarC(slice []int, buscado int) []int {
	var posicion []int
	for i := range slice {
		if slice[i] == buscado {
			posicion = append(posicion, i)
		}
	}
	return posicion
}

func main() {
	var n, buscado int
	fmt.Print("Ingrese el valor de n: ")
	fmt.Scan(&n)

	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(10)+1)
	}

	fmt.Print("Ingrese el numero a buscar: ")
	fmt.Scan(&buscado)

	fmt.Println(slice)
	fmt.Println("La posicion del primer valor buscado es: ", buscar(slice, buscado))

	fmt.Println("los indices son: ", buscarC(slice, buscado))

}
