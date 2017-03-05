package main

import (
  "fmt"
)

func main() {
  // 1

  c1 := make(chan int)
  // c1 <- 1    // this is blocking
  go func() {  // solution
    c1 <- 1 
  }()
  fmt.Println(<-c1)


  // 2

  c2 := make(chan int)

  go func() {
    for i := 0; i < 10; i++ {
      c2 <- i
    }
    close(c2) // solution
  }()

  // fmt.Println(<-c2) // only prints 0
  for n := range c2 {  // solution
    fmt.Println(n)
  }
}