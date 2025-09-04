package main
import "fmt"

func quickSort(arr[]int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := len(arr) / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
		arr[left], arr[i] = arr[i], arr[left]
		left ++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
    quickSort(arr[left+1:])
	return arr
}

func main () {
nums := []int{5, 2, 7, 1}

	fmt.Println ("До сортировки:", nums)
quickSort(nums)
	fmt.Println("После сортировки", nums)
}