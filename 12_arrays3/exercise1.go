package main

import "fmt"

func main() {
	enero := []float64{1.29, 9.99, 22.50, 4.55, 7.35, 6.49}
	febrero := make([]float64, len(enero))

	copy(febrero, enero)

	febrero[2] = 100.99

	fmt.Println(enero)
	fmt.Println(febrero)
}
