package main

import "fmt"

func main() {
	greet("Jane", "Doe") //args
}

func greet(fname, lname string) { //params
	fmt.Println(fname, lname)
}
