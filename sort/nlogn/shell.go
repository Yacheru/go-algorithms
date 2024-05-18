package main

import "fmt"

func main() {
	data := []int{9, 8, 3, 7, 5, 6, 4, 1, 2}
	size := len(data)
	arr := shellSort(data, size)

	fmt.Println(arr)
}

func shellSort(arr []int, size int) []int {
	interval := size / 2

	for interval = size / 2; interval > 0; interval /= 2 {
		for i := interval; i < size; i++ {
			temp := arr[i]
			var j int
			for j = i; j >= interval && arr[j-interval] > temp; j -= interval {
				arr[j] = arr[j-interval]
			}
			arr[j] = temp
		}
	}
	return arr
}
