package main

import "fmt"

func main() {
	arr := []int{48, 50, 20, 23, 7, 24, 5, 49, 34, 9, 26, 4, 44, 21, 32,
		46, 30, 29, 41, 39, 19, 36, 13, 35, 33, 42, 22, 37, 12,
		38, 2, 31, 16, 10, 6, 17, 40, 8, 25, 43, 47, 3, 1, 11, 27,
		45, 18, 15, 14, 28, 52, 51, 55, 53, 54, 56, 57, 60, 58, 59}
	sortArr := selectionSort(arr)

	fmt.Println(sortArr)
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
