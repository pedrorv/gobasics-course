package main

import "fmt"

func main() {
  a := 42

  fmt.Println(a)
  fmt.Println(&a)

  var b *int = &a 

  fmt.Println(b)
  fmt.Println(*b)

  *b = 84
  fmt.Println(a)

  x := 5
  zero(x)
  fmt.Println(x) // x = 5

  zeroPointer(&x)
  fmt.Println(x) // x = 0  
}

func zero(x int) {
  x = 0
}

func zeroPointer(x *int) {
  *x = 0
}