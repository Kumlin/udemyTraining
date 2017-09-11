package main

import "fmt"

func main() {
	switch "Mike" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi", "Mike":
		fmt.Println("Wassup Medhi or Mike")
		fallthrough
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}
