package main

import "net/http"

// Este subpaquete tiene todo lo necesario para convertir nuestro desarrollo en un servidor web

func main() {

	http.HandleFunc("/", home) // Cuando en mi página vaya a / ejecuta la función home
	//http.HandleFunc("/login", login) // Otro ejemplo

	http.ListenAndServe(":3000", nil) // Le pasamos nil porque este parámetro no lo vamos a usar por ahora

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
