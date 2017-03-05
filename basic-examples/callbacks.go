package main

import "fmt"

func main() {
  visit([]int{1, 2, 3, 4}, printNumber)

  xs := filter([]int{1, 2, 3, 4}, filterCondition)
  fmt.Println(xs)
}

func visit(numbers []int, callback func(int)) {
  for _, n := range numbers {
    callback(n)
  }
}

func printNumber(number int) {
  fmt.Println(number)
}

func filter(numbers []int, callback func(int) bool) []int {
  xs := []int{}

  for _, n := range numbers {
    if callback(n) {
      xs = append(xs, n)
    }
  }

  return xs
}

func filterCondition(number int) bool {
  return number > 1
}