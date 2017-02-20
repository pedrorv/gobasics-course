package main

import "fmt"

func main() {
  
  // 1
  fmt.Println("Hello World")

  // 2
  fmt.Println("Hello, my name is Pedro")

  // 3
  fmt.Print("Enter your name: ")
  
  var userName string
  fmt.Scan(&userName)
  fmt.Printf("Hello, %v! \n", userName)

  // 4
  fmt.Print("Enter a small number: ")
  var smallNumber int
  fmt.Scan(&smallNumber)

  fmt.Print("Enter a larger number: ")
  var largerNumber int
  fmt.Scan(&largerNumber)

  fmt.Printf(
    "The remainder of %v/%v is %v. \n", 
    largerNumber, 
    smallNumber,
    largerNumber % smallNumber)

  // 5
  for i := 0; i <= 100; i++ {
    if i % 2 == 0 {
      fmt.Println(i)
    }
  }

  // 6
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 && i % 5 == 0 {
      fmt.Println("FizzBuzz")
    } else if i % 3 == 0 {
      fmt.Println("Fizz")
    } else if i % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }

  // 7
  sum := 0
  for i := 0; i < 1000; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      sum += i
    }
  }
  fmt.Println(sum)

}