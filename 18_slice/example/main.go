package main

import "fmt"

func main() {
	mySlice := make([]string, 5, 10)

	mySlice = append(mySlice, "this")
	mySlice = append(mySlice, "that")

	fmt.Println(len(mySlice))
	fmt.Println(mySlice)
}
