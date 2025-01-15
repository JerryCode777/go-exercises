package main

import (
	"fmt"
	"math/rand"
	"time"
)

const pistaLongitud = 100 // Longitud de la pista (100 metros)

func generarCarrera(nombre string, fila int, c chan string) {
	pos := 0 // Posición inicial del corredor
	for pos < pistaLongitud {
		// Dibujar pista
		fmt.Printf("\033[%d;1H%s |%s%s| %d m\n", fila, nombre, string(repeat('=', pos)), string(repeat(' ', pistaLongitud-pos)), pos)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(150)+50)) // Pausa aleatoria para más velocidad
		pos += rand.Intn(4)                                             // Avance aleatorio (0, 1, 2 o 3 metros por iteración)
	}
	// Indicar que terminó la carrera
	c <- nombre
}

func repeat(char rune, n int) []rune {
	result := make([]rune, n)
	for i := range result {
		result[i] = char
	}
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Semilla para la aleatoriedad

	// Limpiar pantalla
	fmt.Print("\033[2J") // comando apra limpiar pantalla
	fmt.Println("¡Comienza la carrera de goroutines!")
	fmt.Printf("Meta: %d metros\n\n", pistaLongitud)

	// Canales para determinar el ganador
	resultado := make(chan string)

	// Iniciar 8 corredores en diferentes filas
	go generarCarrera("G1", 3, resultado)
	go generarCarrera("G2", 4, resultado)
	go generarCarrera("G3", 5, resultado)
	go generarCarrera("G4", 6, resultado)
	go generarCarrera("G5", 7, resultado)
	go generarCarrera("G6", 8, resultado)
	go generarCarrera("G7", 9, resultado)
	go generarCarrera("G8", 10, resultado)

	// Esperar al ganador
	ganador := <-resultado
	fmt.Printf("\033[%d;1H¡El ganador es %s! 🎉\n", 12, ganador)

	// Esperar brevemente antes de finalizar
	time.Sleep(2 * time.Second)
}
