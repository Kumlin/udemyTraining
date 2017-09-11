package main

import "fmt"

func main() {

	friendsName := "something"

	switch {
	case friendsName == "something":
		fmt.Println("Wassup Daniel")
	case (friendsName == "Medhi") || (friendsName == "Mike"):
		fmt.Println("Wassup Medhi or Mike")
		fallthrough
	case friendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}
