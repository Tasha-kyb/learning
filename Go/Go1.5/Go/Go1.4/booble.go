package main

import "fmt"

func main() {
    numbers := []int{-1, -4, 6, 5, 8, 7, -7, 9}
    fmt.Println("До сортировки:", numbers)
    
    bubbleSort(numbers)
    
    fmt.Println("После сортировки:", numbers)
}

func bubbleSort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        for j := 0; j < len(arr)-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Меняем местами элементы
                arr[j], arr[j+1] = arr[j+1], arr[j]
                // 1ая итерация {-4, -1, 5, 6, 7, -7, 8, 9}
                // 2ая итерация {-4, -1, 5, 6, -7, 7, 8, 9}
                // итого 7 итераций
            }
        } 
    }
}
