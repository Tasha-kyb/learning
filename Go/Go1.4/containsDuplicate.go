package main
import "fmt"

func main () {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println ("Результат:", result)
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}