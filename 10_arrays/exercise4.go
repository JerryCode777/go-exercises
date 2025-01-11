package main

import (
	"fmt"
	"math/rand"
)

func histograma(slice [5]int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print("nota: ", i, " => \t")
		for j := 0; j < slice[i]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	var array [2][5]int
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			array[i][j] = rand.Intn(5) + 1
		}
	}
	var frecuencia [5]int
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			for f := 0; f < 5; f++ {
				if array[i][j] == f+1 {
					frecuencia[f]++
				}
			}
		}
	}
	fmt.Println(frecuencia)
	histograma(frecuencia)
}
