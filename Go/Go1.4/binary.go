package main
import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    target := 4
    
    result := binarySearch(nums, target)
    if result == -1 {
        fmt.Println("Элемент не найден")
    } else {
        fmt.Printf("Элемент %d найден на позиции %d", target, result)
    }
}

func binarySearch(arr []int, target int) int {
    low := 0
    high len(arr)-1 // = 4
    
    for low <= high {
        mid := low + (high-low)/2 // = 2, arr[2] = 3
        
        if arr[mid] == target {
            return mid // Нашли элемент
        }
        
        if arr[mid] < target {
            low = mid + 1 // low = 3
        } else {
            high = mid - 1
        }
    }
    
    return -1 // Элемент не найден
}