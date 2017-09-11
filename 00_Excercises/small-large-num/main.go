package main

import (
	"fmt"
	"time"
)

func main() {
	var smallN int
	var bigN int

	fmt.Println("Enter a small and then large number to calculate the remainder")
	time.Sleep(1000)
	fmt.Print("Small number: ")
	fmt.Scan(&smallN)
	fmt.Print("Large number: ")
	fmt.Scan(&bigN)

	if smallN > bigN {
		fmt.Print("Your small number is bigger than your large one!")
	} else {
		remainder := bigN % smallN
		fmt.Print("The remainder of ", bigN, " divided by ", smallN, " equals ", remainder)
	}
}
