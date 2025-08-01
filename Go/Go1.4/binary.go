package main
import "fmt"

func binarySearch(arr []int, target int) int {
    low, high := 0, len(arr)-1
    
    for low <= high {
        mid := low + (high-low)/2
        
        if arr[mid] == target {
            return mid // Нашли элемент
        }
        
        if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    return -1 // Элемент не найден
}

func main() {
    sorted := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
    target := 23
    
    result := binarySearch(sorted, target)
    if result == -1 {
        fmt.Println("Элемент не найден")
    } else {
        fmt.Printf("Элемент %d найден на позиции %d\n", target, result)
    }
}