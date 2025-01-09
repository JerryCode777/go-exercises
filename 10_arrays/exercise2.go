package main

import (
	"fmt"
	"math/rand"
)

func sumaArray(array [11]int) int {
	var suma int
	for i := 0; i < len(array); i++ {
		suma += array[i]
	}
	return suma
}

func tablaFrecuencias(array [11]int) {
	suma := sumaArray(array)
	for i := 0; i < len(array); i++ {
		fmt.Print(i+2, "\t")
		for j := 0; j < array[i]*100/suma; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var nro, dado1, dado2 int
	fmt.Println("Ingrese el numero de lanzamientos(dos dados): ")
	fmt.Scan(&nro)

	var lanzamientos []int
	var frecuencias [11]int
	for i := 0; i < nro; i++ {
		dado1 = rand.Intn(6) + 1
		dado2 = rand.Intn(6) + 1
		lanzamientos = append(lanzamientos, dado1+dado2)

		for j := 2; j <= 12; j++ {
			if j == lanzamientos[i] {
				frecuencias[j-2]++
			}
		}
	}
	fmt.Println(frecuencias)
	tablaFrecuencias(frecuencias)
}
