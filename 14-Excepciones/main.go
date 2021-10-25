package main

import "log"

// El defer es una función que se va a ejecutar sí o sí cuando se detecta que una función se va
// por un return o por un error o por un fin de función

// Panic es una función que nos permite abortar una función mostrando el mensaje
// que indiquemos por la consola, es como forzar un error

// Para tener el control de un panic existe algo llamado recover, y el recover se va a ejecutar
// cuando detecta que hay un panic y va a tomar el control del mensaje que hemos dado

func main() {

	/*archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close() // Se va a ejecutar (siempre) cuando salga de la función

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1) // Cierra la aplicación directamente / Finaliza la ejecución del programa
	}*/

	ejemploPanic()

}

func ejemploPanic() {

	// Después de un panic no hay control alguno que podamos hacer, salvo que coloquemos un defer
	defer func() {
		reco := recover() // Recover me trae el resultado del último panic y si no hubo panic entonces nos devuelve un nil
		if reco != nil {
			// Es lo mismo que un printf sólo que este va a grabar en el archivo de log y nos muestra por consola la fecha y la hora
			// %v hace referencia a variant de otros lenguajes
			log.Fatalf("ocurrió un error que generó Panic \n %v", reco) // Hace un Printf() y además un os.Exit(1)
		}
	}()

	a := 1
	if a == 1 {
		panic("se encontro el valor de 1") // Aborta el programa GO
	}

}
