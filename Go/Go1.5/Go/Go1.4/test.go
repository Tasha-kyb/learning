package main
import (
	"fmt"
	"sort"
)

// Напишите функцию SumArray(arr []int) int, которая возвращает сумму всех чисел в массиве.
// Напишите функцию FindMax(arr []int) int, которая находит максимальное число в массиве.
//Напишите функцию CountPositives(arr []int) int, 
// которая возвращает количество положительных чисел в массиве.
// Напишите функцию FindMin(arr []int) int, которая возвращает минимальное число в массиве.
// Напишите функцию Average(arr []int) float64, которая возвращает среднее арифметическое чисел массива.

func main () {
nums := []int{5, 2, 7, 1}
sort.Ints(nums) // Использует Quick Sort под капотом
fmt.Println(nums) // [1 2 5 7]	
sort.Sort(sort.Reverse(sort.IntSlice(nums)))
fmt.Println(nums)
}
