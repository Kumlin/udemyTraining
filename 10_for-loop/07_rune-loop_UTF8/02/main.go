package main

import "fmt"

func main() {
	for i := 5000; i <= 6140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}