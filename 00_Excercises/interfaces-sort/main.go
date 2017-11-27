package main

import (
	"fmt"
	"sort"
)

type people []string

func (a people) Len() int {
	return len(a)
}

func (a people) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a people) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
}
