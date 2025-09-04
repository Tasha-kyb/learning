package main
import "fmt"

func main () {
	nums := []int{1, 2, 3, 1}
	result := containsDuplicate(nums)
	fmt.Println ("Результат:", result)
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := m[num]; exist {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}