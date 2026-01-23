// Создай структуру Car с полями:
//Brand (марка машины, строка)
//Year (год выпуска, целое число)
//Напиши функцию NewCar(brand string, year int) Car, которая создаёт и возвращает новый объект Car.
//Напиши метод Description() для Car, который возвращает строку с описанием машины, например:
//"Марка: Toyota, Год выпуска: 2010"
//В функции main() создай несколько машин через NewCar и выведи описание каждой, вызвав метод Description().

package main
import "fmt"

type Car struct {
	Brand string
	Year int
}

func NewCar(brand string, year int) Car {
	return Car {
		Brand: brand,
		Year: year,
	}
}

func (c Car) Description() {
	fmt.Printf("Марка: %s, Год выпуска: %d\n", c.Brand, c.Year)
}

func main() {
	car1 := NewCar("Toyota", 2010)
	car2 := NewCar("BMW", 2018)

	car1.Description()
	car2.Description()
}