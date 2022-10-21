package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubblesort(arr []int) []int {
	temp := 0
	for k := 0; k < len(arr)-1; k++ {
		for i := 0; i < len(arr)-1-k; i++ {
			if arr[i] > arr[i+1] {
				temp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
	return arr
}
func main() {
	var arr = make([]int, 100)
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 100; j++ {
		arr[j] = rand.Intn(100) + 1
	}
	fmt.Println(arr)
	arr = bubblesort(arr)
	fmt.Println(arr)
}
