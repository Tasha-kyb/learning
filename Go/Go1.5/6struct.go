package main

import "fmt"

// Объявление структуры с полями разных типов
type Person struct {Name string; Age int}

func main() {
    // Инициализация структуры
    p := Person{Name: "Иван", Age: 25}
    
    // Вывод значений
    fmt.Printf("Имя: %s\n", p.Name)
    fmt.Printf("Возраст: %d\n", p.Age) 
}