package main

import "fmt"

//findElementsWithSum of k from arr of size
func findElementsWithSum(arr[10] int, combinations[19] int, size int, k int, addValue int, l int, m int) int {
	var num = 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		var p int = 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}

	for i := l; i < size; i++ {
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue + arr[i], l, m + 1)
		l = l + 1
	}
	return num
}

func main() {
	var arr = [10]int{1,4,2,3,6,7,8,4,5,7}
	var addedSum = 18
	var combinations [19]int

	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}