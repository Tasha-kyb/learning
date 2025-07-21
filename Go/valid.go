package main
import "fmt"

func main () {
	s := "([)]"
	result := isValid(s)
	fmt.Println ("Результат:", result)
}

func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{
        ')' : '(',
        '}' : '{',
        ']' : '[',
    }

    for _, currBracket := range s {
        // Если это открывающая скобка - добавляем в стек
        if currBracket == '(' || currBracket == '{' || currBracket == '[' {
            stack = append(stack, currBracket)
            continue
        }

        // Если это закрывающая скобка
        if len(stack) == 0 {
            return false // Нет открывающей скобки для закрытия
        }

        // Проверяем соответствие
		lastOpen := stack[len(stack)-1]
        if pairs[currBracket] != lastOpen {
            return false
        }

		// Берём последнюю открывающую скобку
        stack = stack[:len(stack)-1] // Удаляем её из стека
    }

    // Если стек пуст - все скобки правильно закрыты
    return len(stack) == 0
}