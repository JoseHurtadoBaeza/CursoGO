package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int

var resultado int
var leyenda string

func main() {

	fmt.Println("Ingrese número 1: ")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese número 2: ")
	fmt.Scanln(&numero2)

	fmt.Println("Descripcion : ")

	scanner := bufio.NewScanner(os.Stdin)

	// Controlamos que lo que se ha ingresado es un valor y no dió error
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	resultado = numero1 + numero2

	fmt.Println(leyenda, ":", resultado)

}
