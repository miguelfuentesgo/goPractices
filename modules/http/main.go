package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Run() {

	// HTTP Server

	//Router definitions
	fmt.Println("Http server starting...")

	fmt.Println("Defining routes...")
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)

	fmt.Println("Starting HTTP server...")
	fmt.Println("HTTP server started successfully  and listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("ERROR starting HTTP server", err)
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a mi servidor en Go!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	aboutInfo := map[string]string{
		"author":  "Miguel Fuentes",
		"message": "Este es un servidor creado con Go",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(aboutInfo)
}
