//Создайте структуру Book с полями:
//Title (string) — название книги
//Author (string) — автор книги
//Pages (int) — количество страниц
//Напишите функцию-конструктор NewBook(title, author string, pages int) Book.
//Реализуйте функцию PrintBook(b Book), которая выводит информацию в формате:
//"Книга: <Title>, автор: <Author>, страниц: <Pages>"
//Реализуйте метод Summary() string для Book, который возвращает строку в формате:
//"<Title> — книга автора <Author>, содержит <Pages> страниц"

package main
import "fmt"

type Book struct {
	Title string
	Author string
	Pages int
}

func NewBook(title string, author string, pages int) Book {
	return Book {
	Title : title,
	Author : author,
	Pages : pages,
	}
}

func PrintBook(b Book) {
	fmt.Printf("Книга: %s, автор: %s, страниц: %d\n", b.Title, b.Author, b.Pages)
}

func (b Book) Summary() string {
	return fmt.Sprintf("%s - книга автора %s, содержит %d страниц", b.Title, b.Author, b.Pages)
}

func main() {
	book1 := NewBook("Евгений Онегин", "Александр Пушкин", 300)
	book2 := NewBook("Мертвые души", "Николай Гоголь", 550)

	PrintBook(book1)
	PrintBook(book2)

	fmt.Println(book1.Summary())
	fmt.Println(book2.Summary())
}