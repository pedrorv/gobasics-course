package main

import "fmt"

func main() {
  fmt.Println(average(43, 56, 87, 12, 45, 57))

  data := []float64{43, 56, 87, 12, 45, 57}
  // Variadic arguments
  n := average(data...)
  fmt.Println(n)
}


// Variadic parameters
func average(sf ...float64) float64 {
  // sf is a slice of float64

  var total float64 // same as total := 0.0

  // range gives the index and the value
  // using _ because the index is not used
  for _, v := range sf {
    total += v
  }
  return total / float64(len(sf))
}