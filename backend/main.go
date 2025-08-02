package main

import (
	"fmt"
	"net/http"
)

// Handler para la ruta "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola mundo desde Go y Docker! 🚀")
}

func main() {
	// Asigna handler a la ruta raíz "/"
	http.HandleFunc("/", helloHandler)

	// Puedes agregar más rutas:
	// http.HandleFunc("/ping", pingHandler)

	fmt.Println("Servidor Go escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
