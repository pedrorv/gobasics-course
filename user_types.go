package main

import "fmt"

type foo int

func main() {
  var answer foo
  answer = 42
  fmt.Printf("%T %v \n", answer, answer)

  var answerInt int
  answerInt = 42
  fmt.Printf("%T %v \n", answerInt, answerInt)

  // this doesn't work
  // fmt.Println(answer + answerInt)

  // this conversion works
  fmt.Println(int(answer) + answerInt)
}