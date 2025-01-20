package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	fmt.Println("Ingrese el valor de n")
	fmt.Scan(&n)

	fmt.Println("Ingrese los valores del arreglo: ")
	array := ingresar(n)
	modificar(array)
	imprimir(array)

}

func ingresar(n int) []int {
	arreglo := make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&arreglo[i]); err != nil {
			log.Fatal(err)
		}
	}
	return arreglo
}
func modificar(arreglo []int) []int {
	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = arreglo[i] * 2
	}
	return arreglo
}

func imprimir(arreglo []int) {
	fmt.Print("[ ")
	for _, i := range arreglo {
		fmt.Print(i,i == len(arreglo) ? "":",")
	}
	fmt.Print("]")
	fmt.Println()
}
