package main
import "fmt"

type Person struct {
	Name string
	Age int
}

func NewPerson(name string, age int) Person {
	return Person {
		Name: name,
		Age: age,
	}
}

func (p Person) Greet() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", p.Name, p.Age)
}

func main() {
	ps1 := NewPerson("Андрей", 30)
	ps2 := NewPerson("Марина", 25) 

	ps1.Greet()
	ps2.Greet()
}
