package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado:", areaCuadrado)
}
