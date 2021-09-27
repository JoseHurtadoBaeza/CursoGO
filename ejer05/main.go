package main

import "fmt"

func main() {

	/*for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}*/

	/*numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break // Para romper bucles
		}
	}*/

	/*var i = 0
	for i < 10 {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			i = i * 2
			continue // Manda el control de la ejecución nuevamente al inicio del ciclo for, no sigue, no continua
		}
		fmt.Printf("	Pasó por aquí \n")
		i++

	}*/

	var i int = 0

RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a Rutina sumando 2 a i")
			/*Esto es el como el "continue" visto antes, con la diferencia de que
			en lugar de ir al principio del bucle, va a ir a donde nosotros le digamos*/
			goto RUTINA

		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}

}
