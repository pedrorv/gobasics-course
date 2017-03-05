package main

import "fmt"

func main() {
  // 1
  fmt.Println(half(1))
  fmt.Println(half(2))

  // 2
  {
    halfFuncExpression := func(n int) (int, bool) {
      return n/2, n%2 == 0
    }

    fmt.Println(halfFuncExpression(1))
    fmt.Println(halfFuncExpression(2))
  }

  // 3
  fmt.Println(greatest(10, 2, 1, 17, 5, 12))

  // 4
  // (true && false) || (false && true) || !(false && false)
  // false || false || true == true

  // 5
  foo := func (n ...int) int {
    return 42
  }
  foo()
}

func half(number int) (int, bool) {
  return number/2, number%2 == 0
}

func greatest(numbers ...int) int {
  var x int
  for i, n := range numbers {
    if i == 0 {
      x = n
    } else {
      if n > x {
        x = n
      }
    }
  }
  return x 
}
