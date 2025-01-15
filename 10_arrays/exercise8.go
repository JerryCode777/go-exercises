package main

import (
	"fmt"
	"log"
)

func main() {
	//codigo generador de matriz identidada
	var rows, cols int
	fmt.Println("Ingrese las dimensiones de la matriz: ")
	fmt.Print("Numero de filas: ")
	if _, err := fmt.Scan(&rows); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Numero de columnas: ")
	if _, err := fmt.Scan(&cols); err != nil {
		log.Fatal(err)
	}
	mostrarMatriz(generarIdentidad(rows, cols))

}

func generarIdentidad(rows, cols int) [][]int {
	matriz := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matriz[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if i == j {
				matriz[i][j] = 1
			} else {
				matriz[i][j] = 0
			}
		}
	}
	return matriz
}

func mostrarMatriz(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
