//Счётчик посещений (VisitCounter)
//Создайте структуру VisitCounter, которая содержит поле Visits (int).
//Добавьте метод Increment() для увеличения счётчика на 1.
//Добавьте метод Reset() для сброса счётчика в 0.
//Добавьте метод String() string, который возвращает строку "Посещений: <Visits>".
//В main продемонстрируйте создание VisitCounter, увеличение счётчика несколько раз, вывод значения и сброс.

package main
import "fmt"

type VisitCounter struct {
	Visits int
}

func (v *VisitCounter) Increment() {
	v.Visits++
}

func (v *VisitCounter) Reset() {
	v.Visits = 0
}

func (v VisitCounter) String() string {
	return fmt.Sprintf("Посещений: %d", v.Visits)
}

func main() {
	count := VisitCounter{}
	count.Increment()
	count.Increment()
	fmt.Println(count.String())

	count.Reset()
	fmt.Println(count.String())
}