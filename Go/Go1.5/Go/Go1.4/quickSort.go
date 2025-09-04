package main
import (
	"fmt"
	"sort"
)

var originalNums = []int{5, 2, 7, 1}

func main () {
nums := make([]int, len(originalNums))
copy(nums, originalNums)

sort.Ints(nums)
	fmt.Println(nums)

sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)

	fmt.Println("Оригинал:", originalNums)
}