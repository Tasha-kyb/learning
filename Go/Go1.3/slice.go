package main
import "fmt"

func main () {
	slice := []int {1,2,3,4,5}
	for i := len(slice)-1; i>=0; i -- {
		fmt.Print (slice[i])
	}
}