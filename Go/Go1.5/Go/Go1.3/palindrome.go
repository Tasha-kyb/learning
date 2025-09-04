package main
import "fmt"


func main () {
	x := 121
	result := palindrome (x)
		fmt.Println ("результат:", result)
}

func palindrome (x int) bool {
		if x < 0 {
			return false
		}
		if x == 0 {
			return true
		}

		reversed := 0
		original := x
		
		for x > 0 {
			digit := x%10
			reversed = reversed*10 + digit
			x/=10
		}
		return original == reversed
}