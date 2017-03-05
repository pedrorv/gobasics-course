package main

import (
  "fmt"
  "math"
)

// Shapes types

type Square struct {
  side float64
}

type Circle struct {
  radius float64
}

// Shape interface 

type Shape interface {
  area() float64
}

// Shapes methods

func (z Square) area() float64 {
  return z.side * z.side
}

func (c Circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

// Shape method accepts any shape that implements
// the Shape interface

func info(z Shape) {
  fmt.Println(z)
  fmt.Println(z.area())
}


func main() {
  s := Square{10}
  info(s)

  c := Circle{5}
  info(c)
}