package main

import "fmt"

func main() {
  fmt.Printf("%d! = %d \n", 1, factorial(1))
  fmt.Printf("%d! = %d \n", 2, factorial(2))
  fmt.Printf("%d! = %d \n", 3, factorial(3))
  fmt.Printf("%d! = %d \n", 4, factorial(4))
  fmt.Printf("%d! = %d \n", 5, factorial(5))
}

func factorial(n int) int {
  if n < 1 {
    return 1
  }

  return n * factorial(n - 1)
}