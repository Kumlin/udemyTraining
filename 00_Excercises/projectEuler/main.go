package main

import "fmt"

func main() {
	var answer int
	var tally int

	i := 1
	for {
		fmt.Println(i)
		if isBouncy(i) {
			tally++
		}

		if float64(tally)/float64(i) == 0.99 {
			answer = i
			break
		}
		i++
	}
	println("First bouncy number to have 99% bouncies is: ", answer)
}

//takes a number and stores each digit in a slice
//numbers stored rightmost to leftmost digit
func digitSlicer(number int) []int {
	var digits []int
	for number != 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	//fmt.Println(digits)
	return digits
}

//find if a number's digits are increasing and decreasing in size
func isBouncy(number int) bool {
	slices := digitSlicer(number)
	var lastDigit = slices[0]
	increasing := false
	decreasing := false

	for i := 0; i != len(slices); i++ {
		if slices[i] > lastDigit {
			increasing = true
		} else if slices[i] < lastDigit {
			decreasing = true
		}
		// fmt.Println("iteration: ", i, " currentDigit: ", slices[i], " last digit: ", lastDigit)
		lastDigit = slices[i]
	}

	return increasing && decreasing
}

//find a number with 99% bouncy numbers
//increment down until num < 99%
//solution is increment one up
//problem #112
