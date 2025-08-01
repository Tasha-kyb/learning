package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет, Мир!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Текущее время: %s", time.Now().Format("15:04:05"))
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/time", timeHandler)
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "О нас")
    })
    fmt.Println("Сервер запущен: http://localhost:8080/time")
    http.ListenAndServe(":8080", nil)
}
