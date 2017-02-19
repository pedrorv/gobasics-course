package main

import "fmt"

func main() {
  for i:= 48; i <= 126; i++ {
    fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
  }
}