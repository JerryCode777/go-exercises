package main

import (
	"fmt"
	"math/rand"
)

func lanzamiento(dados int) []int {
	var lanzamiento []int
	for i := 0; i < dados; i++ {
		lanzamiento = append(lanzamiento, rand.Intn(6)+1)
	}
	return lanzamiento
}

func puntajes(jugadas [][]int) []int {
	var puntajes []int
	var suma int
	for i := range jugadas {
		suma = 0
		for j := range jugadas[i] {
			suma += jugadas[i][j]
		}
		puntajes = append(puntajes, suma)
	}
	return puntajes
}
func histograma(slice []int) {
	for i := range slice {
		fmt.Print("jugador ", i+1, " : ")
		for j := 0; j < slice[i]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	var n, m int
	fmt.Print("Ingrese el valor del numero de dados: ")
	fmt.Scan(&n)
	fmt.Print("Ingrese el numero de concursantes: ")
	fmt.Scan(&m)

	var jugadas [][]int
	for i := 0; i < m; i++ {
		jugadas = append(jugadas, lanzamiento(n))
	}

	puntajes := puntajes(jugadas)
	histograma(puntajes)
	fmt.Println("El ganador es el jugador: ", ganador(puntajes))

}

func ganador(slice []int) int {
	var maximo, indice int
	for i := range slice {
		if maximo < slice[i] {
			maximo = slice[i]
			indice = i
		}
	}
	return indice + 1
}
