package main
import "fmt"

func main () {
	nums := []int{-1, -4, 6, 5, 8, 7,7, 9}
	result := containsDuplicate (nums)
	fmt.Println ("Результат:", result)
}

func containsDuplicate (nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		if _, exist := m[num]; exist {
			return true
		}
		m[num] = true
	}
	return false
}