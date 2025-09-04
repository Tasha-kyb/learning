package main
import "fmt"

func reverse(x int) int {
    reversed := 0
    for x != 0 {
        reversed = reversed*10 + x%10
        x = x / 10
    }
    return reversed
}

func main() {
    fmt.Println(reverse(123))
    fmt.Println(reverse(-123))
    fmt.Println(reverse(120))
    fmt.Println(reverse(0))
}