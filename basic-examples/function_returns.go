package main

import "fmt"

func main() {
  fmt.Println(greet("John", "Doe"))
  fmt.Println(greetNamedReturn("John", "Doe"))
  fmt.Println(greetMultipleReturns("John", "Doe"))
}

func greet(fname, lname string) string {
  return fmt.Sprint(fname, " ", lname)
}

func greetNamedReturn(fname, lname string) (s string) {
  s = fmt.Sprint(fname, " ", lname)
  return
}

func greetMultipleReturns(fname, lname string) (string, string) {
  return fmt.Sprint(fname, " ", lname), fmt.Sprint(fname, " ", lname)
}