package main
import "fmt"

type Node struct {
	Name string
	Age int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append (name string, age int) {
	current := list.Head 
	if list.Head == nil {
		list.Head = &Node{Name: name, Age: age}
		return
	}
	for current.Next != nil {
		current = current.Next
	} 
	current.Next = &Node{Name: name, Age: age}
}

func (list *LinkedList) Print() {
	current := list.Head
    for current != nil {
        fmt.Printf("%s(%d)", current.Name, current.Age)
        if current.Next != nil {
            fmt.Print(" -> ")
        }
        current = current.Next
    }
	fmt.Println()
}

func (list *LinkedList) Delete (name string) { 
	if list.Head == nil {
		return
	}
	if list.Head.Name == name {
		list.Head = list.Head.Next
		return
	}
	current := list.Head
	for current.Next != nil{
		if current.Next.Name == name {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func main() {
	list := &LinkedList{}
	list.Append("Оля", 31)
	list.Append("Маша", 29)
		fmt.Println("Текущий список:")
	list.Print()
	list.Append("Юля", 40)
	list.Append("Катя", 37)
		fmt.Println("После добавления:")
	list.Print()
	list.Delete("Катя")
		fmt.Println("После удаления:")
	list.Print()
}