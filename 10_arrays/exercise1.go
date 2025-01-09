package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//lanzamiento de un dado simulacion
	var nroLanzamientos int
	fmt.Print("Ingrese numero de lanzamientos del dado: ")
	fmt.Scan(&nroLanzamientos)

	var lanzamiento []int
	var frecuencias [6]int
	for i := 0; i < nroLanzamientos; i++ {
		lanzamiento = append(lanzamiento, rand.Intn(6)+1)
		switch lanzamiento[i] {
		case 1:
			frecuencias[0]++
		case 2:
			frecuencias[1]++
		case 3:
			frecuencias[2]++
		case 4:
			frecuencias[3]++
		case 5:
			frecuencias[4]++
		case 6:
			frecuencias[5]++
		}
	}
	fmt.Println(lanzamiento)
	fmt.Println(frecuencias)
}
