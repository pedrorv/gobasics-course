package main

import (
  "fmt"
)

func main() {
  fc := factorial(5)
  for n := range fc {
    fmt.Println(n)
  }
}

func factorial(n int) chan int {
  out := make(chan int, n)
  go func() {
    total := 1
    for i := n; i > 0; i-- {
      total *= i
    }
    out <- total
    close(out)
  }()
  return out
}
