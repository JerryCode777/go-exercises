package main

import (
	"fmt"
	"log"
)

func main() {
	var rows, cols int
	fmt.Println("Ingrese las dimensiones de las matrices")
	fmt.Print("Ingrese numero de filas: ")
	if _, err := fmt.Scan(&rows); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Ingrese numero de columnas: ")
	if _, err := fmt.Scan(&cols); err != nil {
		log.Fatal(err)
	}

	A := readMatriz(rows, cols)
	fmt.Println("Matriz A: ")
	mostrarMatriz(A)
	B := readMatriz(rows, cols)
	fmt.Println("Matriz B: ")
	mostrarMatriz(B)

	fmt.Println("Resultados: ")
	fmt.Println("==================================")

	fmt.Println("A+B")
	suma, err := sumaMatriz(A, B)
	if err != nil {
		log.Fatal(err)
	}
	mostrarMatriz(suma)

	fmt.Println("A-B")
	resta, err := restaMatriz(A, B)
	if err != nil {
		log.Fatal(err)
	}
	mostrarMatriz(resta)

	fmt.Println("A*B")
	producto, err := productoMatriz(A, B)
	if err != nil {
		log.Fatal(err)
	}
	mostrarMatriz(producto)

	fmt.Println("==================================")

}

// lee los elementos de una matriz de dimensiones rows x cols desde la entrada estandar
func readMatriz(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Printf("Elemento [%d][%d]: ", i, j)
			if _, err := fmt.Scan(&matrix[i][j]); err != nil {
				log.Fatalf("Error al leer el elemento [%d] [%d]: %v", i, j, err)
			}
		}
	}
	return matrix
}

func sameDimensions(A, B [][]int) bool {
	if len(A) == 0 || len(B) == 0 || len(A) != len(B) {
		return false
	}
	for i := range A {
		if len(A[i]) != len(B[i]) {
			return false
		}
	}
	return true
}

func sumaMatriz(array1, array2 [][]int) ([][]int, error) {
	if !sameDimensions(array1, array2) {
		return nil, fmt.Errorf("Las matrices deben tener mismas dimensiones")
	}
	m := len(array1)
	n := len(array1[0])

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = array1[i][j] + array2[i][j]
		}
	}

	return result, nil
}

func restaMatriz(array1, array2 [][]int) ([][]int, error) {
	if !sameDimensions(array1, array2) {
		return nil, fmt.Errorf("Las matrices deben tener mismas dimensiones")
	}
	m := len(array1)
	n := len(array1[0])

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = array1[i][j] - array2[i][j]
		}
	}
	return result, nil
}

func productoMatriz(array1, array2 [][]int) ([][]int, error) {
	if !sameDimensions(array1, array2) {
		return nil, fmt.Errorf("Las matrices deben tener mismas dimensiones")
	}
	m := len(array1)
	n := len(array1[0])

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			result[i][j] = array1[i][j] * array2[i][j]
		}
	}
	return result, nil
}

func mostrarMatriz(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
