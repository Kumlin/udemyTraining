package main

import (
  "fmt"
  "sync"
  "runtime"
)

var waitGroup sync.WaitGroup

func init() {
  runtime.GOMAXPROCS(runtime.NumCPU())
}
''
func main () {
  waitGroup.Add(2)
  go mishMash("Foo: ")
  go mishMash("Bar: ")
  waitGroup.Wait()
}

func mishMash(threadName string) {
  for i := 0; i < 1000; i++ {
    fmt.Println(threadName, " ", i)
  }
  waitGroup.Done()
}
