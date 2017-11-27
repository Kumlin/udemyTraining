package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	numberSet := []int{1, 2, 3, 4, 5, 4, 3}

	for i := len(numberSet); i < 100000; i++ {
		numberSet = append(numberSet, rand.Intn(15))
	}

	inbound := sender(numberSet)

	chan1 := factorial(inbound)
	chan2 := factorial(inbound)
	chan3 := factorial(inbound)
	chan4 := factorial(inbound)
	chan5 := factorial(inbound)
	chan6 := factorial(inbound)
	chan7 := factorial(inbound)
	chan8 := factorial(inbound)
	chan9 := factorial(inbound)
	channelSet := []chan int{chan1, chan2, chan3, chan4, chan5, chan6, chan7, chan8, chan9}

	for n := range merge(channelSet) {
		fmt.Println(n)
	}

	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)
}

func merge(channels []chan int) chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, c := range channels {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

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
