package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
