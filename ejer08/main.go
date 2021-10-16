// ARREGLOS ESTÁTICOS Y SLICES
package main

import "fmt"

func main() {

	//matriz := []int{2, 5, 4}
	//fmt.Println(matriz)
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:] // Desde la posición 3 hasta la última
	//porcion := elementos[2:4] // De la 2 a la 4

	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20) // Me reserva un espacio en memoria de hasta 20 elementos, pero inicialmente me va a crear un vector de 5 elementos
	// El printf nos permite mezclar texto con datos
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0) // Me crea un slice vacío, tendremos que hacer appends

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
