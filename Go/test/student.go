//Создай структуру Student с полями:Name (имя студента, строка), Grade (оценка, целое число)
//Напиши функцию NewStudent(name string, grade int) Student, которая создаёт и возвращает новый объект Student.
//Напиши функцию PrintStudent(s Student), которая принимает объект Student и выводит информацию в формате:
//"Студент: <Name>, оценка: <Grade>"
//Напиши метод Info() для Student, который возвращает строку с информацией в том же формате.
//В функции main() создай несколько студентов через NewStudent. Для каждого:
//вызови функцию PrintStudent
//выведи строку, возвращаемую методом Info() (через fmt.Println)

package main
import "fmt"

type Student struct {
	Name string
	Grade int
}

func NewStudent (name string, grade int) Student {
	return Student {
		Name: name,
		Grade: grade,
	}
}

func PrintStudent(s Student) {
	fmt.Printf("%s, оценка: %d\n", s.Name, s.Grade)
}

func (s Student) PrintStudent() string{
	return fmt.Sprintf("%s, оценка: %d", s.Name, s.Grade)
}

func main() {
	s1 := NewStudent("Аня", 9)
	s2 := NewStudent("Маша", 6)

	PrintStudent(s1)
	PrintStudent(s2)

	fmt.Println(s1.PrintStudent())
	fmt.Println(s2.PrintStudent())
}
