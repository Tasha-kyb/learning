// Задание 1. Реализуй бинарный поиск

package main
import "fmt"

func main () {
    nums := []int{1, 2, 3, 4, 5}
    target := 5
	result := binarySearch(nums, target)
    	if result == -1 {
        	fmt.Println("Элемент не найден")
    	} else {
        	fmt.Printf("Элемент %d найден на позиции %d\n", target, result)
    	}
}

func binarySearch(arr[]int, target int) int {
    low := 0
    high := len(arr)-1

    for low <= high {
        mid := low + (high - low)/2

        if arr[mid] == target {
        return mid
        }

        if arr[mid] < target {
        low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}