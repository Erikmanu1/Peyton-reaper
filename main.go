package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "Olá, esta é a API Peyton-reaper em Go!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Servidor rodando em http://localhost:8080/hello")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
