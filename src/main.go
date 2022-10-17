package main

import "fmt"

func main() {

	// For condicinal
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	// For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Printf("\n")

	// For forever
	// Tener cuidado. Se va a ejecutar infinitamente ya que no tiene ninguna condición para cerrarse.
	// counterForever := 0
	// for {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	fmt.Printf("\n")

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
