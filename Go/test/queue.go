/* Добавьте в очередь клиентов с id: 101, 102, 103.
Обслужьте (удалите из очереди) одного клиента.
Добавьте еще клиентов — 104 и 105.
Обслужьте двух клиентов.
Выведите список оставшихся в очереди клиентов.*/

package main
import "fmt"

type Queue struct {
	queue []int
}

func (q *Queue) Enqueue(value int) {
	q.queue = append(q.queue, value)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.queue) == 0 {
		return 0, false
	}
	value := q.queue[0]
	q.queue = q.queue[1:]
	return value, true
}

func main() {
	values := &Queue {}
	values.Enqueue(101)
	values.Enqueue(102)
	values.Enqueue(103)
	fmt.Printf("Очередь: %v\n", values.queue)
	values.Dequeue()
	fmt.Printf("Обслужили одного: %v\n", values.queue)
	values.Enqueue(104)
	values.Enqueue(105)
	fmt.Printf("Очередь увеличилась: %v\n", values.queue)
	values.Dequeue()
	values.Dequeue()
	fmt.Printf("Обслужили двоих: %v\n", values.queue)
}

