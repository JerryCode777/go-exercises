package main

import (
	"fmt"
	"math/rand"
)

func histograma(slice [21]int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print("nota: ", i, " => \t")
		for j := 0; j < slice[i]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	//simulacion de las notas de 20 alunmnos
	var alumnos int
	fmt.Println("Ingrese el numero de alumnos: ")
	fmt.Scan(&alumnos)

	var notas []int
	for i := 0; i < alumnos; i++ {
		notas = append(notas, rand.Intn(21))
	}
	fmt.Println(notas)

	var frecuencia [21]int
	for j := 0; j < alumnos; j++ {
		for k := 0; k <= 20; k++ {
			if notas[j] == k {
				frecuencia[k]++
			}
		}
	}
	histograma(frecuencia)
}
