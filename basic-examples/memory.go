package main

import "fmt"

func main() {
  a := 42

  fmt.Println("a - ", a)
  fmt.Println("a's memory address - ", &a)
  fmt.Printf("a's decimal memory address - %d\n", &a)

  const metersToYards float64 = 1.09361

  var meters float64
  fmt.Print("Enter meters: ")
  fmt.Scan(&meters)
  yards := meters * metersToYards
  fmt.Println(meters, " meters is ", yards, " yards.")
}