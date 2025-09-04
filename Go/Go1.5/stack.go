package main
import "fmt"

type Stack struct {
	values []int 
}

func (s *Stack) Push(value int) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.values) == 0 {
		return 0, false
	}
	lastIndex := len(s.values)-1
	value := s.values[lastIndex]

	s.values = s.values[:lastIndex]
	return value, true
}

func main() {
	stack := &Stack{}
	stack.Push(31)
		fmt.Println("Текущий список:", stack.values)
	stack.Push(29)
	stack.Push(11)
		fmt.Println("Список после добавления:", stack.values)
	stack.Pop()
		fmt.Println("Список после удаления:", stack.values)
}