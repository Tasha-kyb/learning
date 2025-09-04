package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
	Message string
	About   string
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/about" {
		fmt.Fprintln(w, h.About)
	} else {
		fmt.Fprintln(w, h.Message)
	}
}

func main() {
	handler := MyHandler{Message: "Привет, Мир!", About: "О нас"}
	
	fmt.Println("Сервер запущен: http://localhost:8080/")
	fmt.Println("Также доступно: http://localhost:8080/about")
	http.ListenAndServe(":8080", handler)
}