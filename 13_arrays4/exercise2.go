package main

import "fmt"

func bubbleSort(arreglo []int) {
	for i := 0; i < len(arreglo)-1; i++ {
		for j := 0; j < len(arreglo)-1-i; j++ {
			if arreglo[j] > arreglo[j+1] {
				arreglo[j], arreglo[j+1] = arreglo[j+1], arreglo[j]
			}

		}
	}
}

func main() {
	//bubble sort
	arreglo := []int{12, 1, 2, 3, 4}
	bubbleSort(arreglo)
	fmt.Println(arreglo)
}
