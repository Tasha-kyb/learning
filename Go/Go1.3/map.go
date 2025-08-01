package main
import "fmt"

func main () {
	Mymap := map[string]int {"один" : 1, "два" : 2,	"три" : 3}
	fmt.Println ("Сумма:", sumMap(Mymap))
}

func sumMap(m map[string]int) int {
    sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}