package main
import "fmt"

type Queue struct {
	values []int
}

func (q *Queue) Enqueue(value int) {
	q.values = append (q.values, value)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.values) == 0 {
		return 0, false
	}

	value := q.values[0]
	q.values = q.values[1:]
	return value, true
}

func main() {
	queue := &Queue{}
	queue.Enqueue(31)
		fmt.Println("Текущий список:", queue.values)
	queue.Enqueue(29)
	queue.Enqueue(11)
		fmt.Println("Список после добавления:", queue.values)
	queue.Dequeue()
		fmt.Println("Список после удаления:", queue.values)
}