package main

import "fmt"

func main() {
	halfer := func(number int) (int, bool) {
		return number / 2, number%2 == 0
	}

	fmt.Println(halfer(1))
	fmt.Println(halfer(2))
}
