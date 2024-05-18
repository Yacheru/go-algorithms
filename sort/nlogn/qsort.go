package main

import "fmt"

func main() {
	sorted := qSort([]int{10, 5, 2, 3, 1, 4, 6, 10, 22, -1, 0, -22, 32, -2})
	fmt.Println(sorted)
}

func qSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // Базовый случай
	} else {
		pivot := arr[0]   // Опорный элемент (Рекурсивный случай)
		var less []int    // Подмассив всех элементов меньше опорного
		var greater []int // Подмассив всех элементов больше опорного

		for _, e := range arr[1:] {
			if e < pivot {
				less = append(less, e)
			} else {
				greater = append(greater, e)
			}
		}
		return append(append(qSort(less), pivot), qSort(greater)...)
	}
}
