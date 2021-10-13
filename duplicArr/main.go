package main

import "fmt"

var arr = [10]int{54, 76, 63, 8, 12, 85, 68, 21, 97, 64}

func main() {
	var c int
	for i := 0; i < len(arr); i++ {
		j := 0
		for j < len(arr) {
			if arr[i] == arr[j] && i != j && i < j {
				fmt.Printf("Yes, array have duplicate elements\n")
				break
			} else {
				c++
			}
			j++
		}
		if j != len(arr) {
			break
		}
	}
	if c == 100 {
		fmt.Println("No, array doesn't have any duplicates")
	}
}
