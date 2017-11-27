package main

import "fmt"

func main() {
	chan1 := incrementor("Foo: ")
	chan2 := incrementor("Bar: ")
	chan3 := puller(chan1)
	chan4 := puller(chan2)
	fmt.Println("Final Counter: ", <-chan3+<-chan4)
}

func incrementor(name string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(name, i)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
