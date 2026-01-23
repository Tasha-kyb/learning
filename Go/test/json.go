package main

import (
    "encoding/json"
    "fmt"
)

// Определяем структуру с экспортируемыми полями и тегами для JSON
type Person struct {
    Name string `json:"name"` // Имя в JSON будет "name"
    Age  int    `json:"age"`  // Возраст в JSON будет "age"
}

func main() {
    // Создаём исходную структуру
    data := Person{Name: "Иван", Age: 30}

    // --- Маршалинг: структура в JSON ---
    jsonBytes, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Ошибка маршалинга:", err)
        return
    }
    jsonStr := string(jsonBytes)
    fmt.Println("Маршалинг (структура → JSON):")
    fmt.Println(jsonStr)

    // --- Демаршалинг: JSON обратно в структуру ---
    var newPerson Person
    err = json.Unmarshal(jsonBytes, &newPerson)
    if err != nil {
        fmt.Println("Ошибка демаршалинга:", err)
        return
    }
    fmt.Println("Демаршалинг (JSON → структура):")
    fmt.Println(newPerson)
}