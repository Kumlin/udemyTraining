package main

import (
	"fmt"
	"net"
)

func main() {
	dialer, err := net.Dial("tcp", "localhost:9609")
	if err != nil {
		panic(err)
	}

	defer dialer.Close()

	fmt.Fprintln(dialer, "I dialed you")
}
