package main

import "fmt"

func selectionSort(arreglo []int) {
	for i := 0; i < len(arreglo); i++ {
		menorInd := i
		for j := i + 1; j < len(arreglo); j++ {
			if arreglo[j] < arreglo[menorInd] {
				menorInd = j
			}
		}
		arreglo[i], arreglo[menorInd] = arreglo[menorInd], arreglo[i]
	}
}

func main() {
	array := []int{1, 32, 2, 32, 43, 342, 12, 32, 4, 3, 5, 654, 8, 5, 4, 34, 5, 5}
	fmt.Println(array)
	selectionSort(array)
	fmt.Println("after selection sort")
	fmt.Println(array)
}
