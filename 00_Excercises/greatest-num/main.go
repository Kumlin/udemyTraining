package main

import "fmt"

func main() {
	numbers := largest(12, 4, 54, 66, 12, 92, 15)

	fmt.Println(numbers)
}

func largest(sf ...float64) float64 {
	var largest float64

	//largest = -1

	for _, v := range sf {
		if v > largest {
			largest = v
		}
	}

	return largest
}
