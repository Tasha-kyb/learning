package main
import "fmt"

func firstUniqChar(s string) (int, rune) {
    // нужно найти первый неповторяющийся символ в строке
    m := make (map[rune]int)
    for _, char := range s {
        m[char] ++
    }

    for i, char := range s {
        if m[char] == 1 {
            return i, char
        }
    }
    return -1, 0
}

func main() {
	index, char := firstUniqChar("leetcode")
	fmt.Println("Индекс:", index, "символ:", string(char))
}