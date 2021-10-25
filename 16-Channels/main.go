package main

import (
	"fmt"
	"time"
)

// Un canal es un espacio de memoria de diálogo en el que van a dialogar distintas
// rutinas y cuando se aloje un valor en ese espacio de memoria, la rutina que está pidiendo
// un valor a cambio, va a actuar en consecuencia. Ese espacio de memoria tiene que ser de un tipo de dato.

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1) // Ejecuta lo que haya en la rutina bucle, pero toma en cuenta que te estoy pasando un canal para dialogar con otras rutinas
	fmt.Println("Llegué hasta aquí")

	msg := <-canal1
	fmt.Println(msg)

}

func bucle(canal chan time.Duration) {
	inicio := time.Now() // Guarda la hora, minuto y segundo del momento en el que entramos a la rutina
	for i := 0; i < 100000000000; i++ {

	}

	final := time.Now()        // Así tenemos la hora de inicio de antes y la final ahora
	canal <- final.Sub(inicio) // Esta función retorna la duración y dicha duración es un tipo de dato, por eso hemos usado Duration antes

}
