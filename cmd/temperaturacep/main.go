package main

import (
	"net/http"

	"github.com/jorgemarinho/temperatura-por-cep/internal/infra/web"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("/clima", web.BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}
