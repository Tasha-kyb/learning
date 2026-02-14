package main
import "fmt"

func GetCriminal() map[string]bool {
		return map[string]bool{
		"Иван": true,
		"Андрей": false,
	}
}

func main() {
	criminal := GetCriminal()
	c, ok := criminal["Андрей"]

	if !ok {
		fmt.Println("Человека нет в базе")
		return
	} 
	
	if c == true {
		fmt.Println("Человек судим")
	} else {
		fmt.Println("Человек НЕ судим")
	}
}