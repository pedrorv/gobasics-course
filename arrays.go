package main

import "fmt"

func main() {
  var x [58]string

  for i := 65; i < 123; i++ {
    x[i-65] = string(i)
  }
  fmt.Println(x)

  var y [256]int 

  for i := 0; i < 256; i++ {
    y[i] = i
  }

  for _, v := range y {
    fmt.Printf("%v - %T - %b\n", v, v, v)
  }
}