package main

import "fmt"

func main() {
	firstNum := 1
	secNum := 2

	fmt.Println(half(firstNum))
	fmt.Println(half(secNum))
}

func half(number int) (int, bool) {
	return number / 2, number%2 == 0
}
