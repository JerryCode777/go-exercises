package main

import (
	"fmt"
	"math/rand"
)

func mostrarArray(array [10]int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func main() {
	var array [10]int
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(20)
	}

	mostrarArray(array)
}
