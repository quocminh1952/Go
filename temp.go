package main

//[2, 10, 112, 1298, 23, 112414]
import (
	"fmt"
)

func count(num int) int {
	var count int
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func countNumber(arr []int, number int) []int {
	result := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if count(arr[i]) == number {
			result = append(result, arr[i])
		}
	}
	return result
}

func main() {
	arr := []int{2, 10, 112, 1298, 23, 112414}
	fmt.Println(countNumber(arr, 4))
}
