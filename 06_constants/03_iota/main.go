package main

import "fmt"

//iota constants
const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
