package main

import "fmt"

var num = []int{2, 10, 20, 50, 60, 25, 20, 15}

func main() {
	res()
}

func res() {
	var sumEv, sumOd int
	for i := 1; i < len(num); i++ {

		if i%2 == 0 {
			sumEv += num[i]
		} else {
			sumOd += num[i]
		}

	}

	if sumOd == sumEv {
		fmt.Println(" Yes\n", "Sum=", sumEv)
	} else {
		fmt.Println(" No\n", "Diff=", uint(sumEv-sumOd))
	}
}
