package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	numberSet := []int{1, 2, 3, 4, 5, 4, 3}

	for i := len(numberSet); i < 100000; i++ {
		numberSet = append(numberSet, rand.Intn(15))
	}

	for n := range factorial(sender(numberSet)) {
		fmt.Println(n)
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)
}

func sender(numbers []int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func factorial(inbound chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range inbound {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()
	return out
}
