//Работа со срезами и методами
//Создайте структуру IntList с полем Items []int.
//Добавьте метод Add(value int) для добавления числа в список.
//Добавьте метод Sum() int, который возвращает сумму всех элементов.
//Добавьте метод Average() float64, который возвращает среднее арифметическое элементов (если список пустой, вернуть 0).
//В main создайте IntList, добавьте несколько чисел и выведите их сумму и среднее.

package main 
import "fmt"

type Interface interface {
	Add(value int)
	Sum() int
	Average() float64
}

type IntList struct {
	Items []int
}

func (i *IntList) Add (value int) {
	i.Items = append(i.Items, value)
}

func (il IntList) Sum() int {
	sum := 0
	for i := 0; i < len(il.Items); i++ {
		sum += il.Items[i]
	}
	return sum
}

func (i IntList) Average() float64 {
	if len(i.Items) == 0 {
		return 0
	}
	return float64(i.Sum()) / float64(len(i.Items))
}

func main() {
	list := IntList{Items: []int {1, 2, 3}}
	list.Add(4)
	fmt.Println(list.Items)

	fmt.Println(list.Sum())
	fmt.Println(list.Average())

}