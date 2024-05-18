package main

import "fmt"

func main() {
	key := 10 // Искомый элемент
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
		21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60} // Требует отсортированный массив
	index, ok := binarySearch(arr, key)

	fmt.Println(index, ok)
}

// binarySearch Возвращает индекс искомого элемента
func binarySearch(arr []int, key int) (index int, find bool) {
	/*
		В low и high хранятся границы той части массива, в которой выполняется поиск
	*/
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2 // Индекс среднего элемента
		guess := arr[mid]       // Значение по индексу среднего элемента
		if guess == key {
			return mid, true // Элемент найдено
		} else if guess > key {
			high = mid - 1 // Половина больше, чем искомый элемент
		} else {
			low = mid + 1 // Половина меньше, чем искомый элемент
		}
	}
	return -1, false // Элемент не найдено
}
