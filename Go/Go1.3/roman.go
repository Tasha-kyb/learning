package main
import "fmt"

func main (){
	s := "MCMXCIV"
	total := romanToInt (s)
	fmt.Println ("Результат:", total)
}

func romanToInt(s string) int {
	m := map[byte]int {
		'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
	}
	result :=0
	prevNum :=0

	for i:= len(s)-1; i>=0; i-- {
		currNum := m[s[i]]
		if currNum < prevNum {
			result -= currNum
		} else {
			result += currNum
		}
		prevNum = currNum
	}
	return result 
	}