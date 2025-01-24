package main

import "fmt"

func insertionSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = key
	}
}

func main() {
	array := []int{2, 3, 43, 12, 32, 423, 423, 132, 32, 4435, 52, 42, 43, 5, 23, 342, 32}
	insertionSort(array)
	fmt.Println("Arrelo ordenado by insertion Sort")
	fmt.Println(array)

}
