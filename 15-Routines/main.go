// ROUTINES DE GO

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLentoo("Jose Hurtado") // Ejecuta la función de manera asíncrona
	fmt.Println("Estoy aquí")
	var x string
	fmt.Scanln(&x) // Va a pedir por teclado que ingrese un valor
}

func miNombreLentoo(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		// De esta forma cada letra tardará un segundo en mostrarse
		time.Sleep(1000 * time.Millisecond) // Esto va a provocar una pausa en la ejecución
		fmt.Println(letra)
	}
}
