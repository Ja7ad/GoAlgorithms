package main

import "fmt"

func findElement(arr[10] int, k int) bool {
	var i int
	for i = 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

func main() {
	var arr = [10]int{1,2,3,4,5,6,3,6,8,9}
	var check = findElement(arr, 10)
	fmt.Println(check)
	var check2 = findElement(arr, 9)
	fmt.Println(check2)
}
