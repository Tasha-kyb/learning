package main
import "fmt"

func moveZeroes(nums []int) {
	zero := 0
	for i := 0; i < len(nums); i ++ {
		if nums [i] != 0 {
			nums[zero], nums[i] = nums[i], nums[zero]
			zero ++
		}
	}
}

func main() {
	nums := []int {0, 1, 4, 7, 0, 8, 0, 8}
	fmt.Println("До сортировки:", nums)
	moveZeroes(nums)
	fmt.Println("После сортировки:", nums)
}