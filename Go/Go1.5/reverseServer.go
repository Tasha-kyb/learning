package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Numbers struct {
    Items []int
}

type ReverseHandler struct{}

func (h ReverseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var numbers Numbers
    json.NewDecoder(r.Body).Decode(&numbers)
    
    for i, j := 0, len(numbers.Items)-1; i < j; i, j = i+1, j-1 {
        numbers.Items[i], numbers.Items[j] = numbers.Items[j], numbers.Items[i]
    }
    
    json.NewEncoder(w).Encode(numbers)
}

func main() {
    handler := ReverseHandler{}
    http.Handle("/", handler)
    
    fmt.Println("Сервер работает: http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}