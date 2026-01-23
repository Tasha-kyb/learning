package main
import "fmt"


func IsPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if i == len(s) - i {
			return true
		}
	}
	return false
}

func main() {
	palindrome := "ABBAA"
	result := IsPalindrome(palindrome)

	fmt.Println(result)

}