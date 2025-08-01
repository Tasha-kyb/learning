package main
import "fmt"

	var nums = []int{2, 7, 11, 15}
	var target = 9

func main () {
	result := twosum (nums, target)
	fmt.Println("результат", result)
}

func twosum (nums []int, target int) []int {
	m := make(map[int]int)
	for i :=0; i < len(nums); i ++ {
		if idx, ok := m[target-nums[i]]; ok {
			return []int{i,idx}
		}
		m[nums[i]] = i
		// {2;0}
	}
	return []int{-1, -1}
}