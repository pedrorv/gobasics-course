package main

import "fmt"

func main() {
  answer := 42
  var language string = "go"
  pi := 3.14
  boolean := true

  fmt.Printf("%v %T \n", answer, answer)
  fmt.Printf("%v %T \n", language, language)
  fmt.Printf("%v %T \n", pi, pi)
  fmt.Printf("%v %T \n", boolean, boolean)
}