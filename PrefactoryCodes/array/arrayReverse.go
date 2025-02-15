package main

import "fmt"

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))
}