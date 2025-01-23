package main

import "fmt"

// binary search
func binarySearch(arreglo []int, valor int) int {
	var alta, media, baja int
	baja = 0
	alta = len(arreglo) - 1

	for baja <= alta {
		media = (alta + baja) / 2
		if arreglo[media] == valor {
			return media
		} else if arreglo[media] > valor {
			alta = media - 1
		} else {
			baja = media + 1
		}
	}
	return -1
}
func main() {
	arreglo := []int{1, 2, 3, 6, 12, 14, 16, 18, 20, 21}
	fmt.Println(binarySearch(arreglo, 14))

}
