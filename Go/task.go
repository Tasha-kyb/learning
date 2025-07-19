package main
import "fmt"

func main () {
    // Задание 1. Как объявить переменные в Go
    /*var c int
    var d int = 9
    
    e := 100 */

    // Задание 2: функция для суммы двух чисел
    fmt.Println ("2 задание:")
    a, b := 5, 10
    result := sum (a, b)
    fmt.Println (result)

    // Задание 3: цикл for от 1 до 5
    fmt.Println ("\n3 задание:")
    for i := 1; i <= 5; i++ {
    fmt.Print (i)
    }

    // Задание 4: напишите условие, которое проверяет, является ли число положительным
    fmt.Println ("\n4 задание:")
    num := -140
    if isPositive(num) {
        fmt.Println ("положительное")
    } else {
        fmt.Println ("отрицательное")
    }

    // Задание 5. Объявите массив из 3 целых чисел
    fmt.Println ("\n5 задание:")
    arr := [3] int{1, 2, 3}
    fmt.Println ("массив:", arr)

    // Задание 6. Создайте слайс чисел и добавьте к нему элемент
    fmt.Println ("\n6 задание:")
    slice := []int{1, 2, 3}
    fmt.Println ("слайс до:", slice)
    slice = append (slice, 4)
    fmt.Println ("слайс после:", slice)

    // Задание 7. Создайте карту, где ключ - строка, значение - число
    fmt.Println ("\n7 задание:")
    myMap := map[string]int {
        "один":1, 
        "два":2, 
        "три":3}
    fmt.Println ("карта",myMap)

    // Задание 8. Как проверить, существует ли ключ в карте?
    fmt.Println ("\n8 задание:")
    keytoCheck := "два"
    value, ok := myMap [keytoCheck]
    if ok {
        fmt.Println ("ключ", keytoCheck, "есть в карте, его значение", value)
    } else {
        fmt.Println ("ключа", keytoCheck, "нет в карте")
    }

}

func sum (a int, b int) int {
    sum := a + b
    return sum
}

func isPositive (num int)bool {
    return num > 0
}