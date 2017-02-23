package main

import "fmt"

type Person struct {
  First string
  Last  string
  Age   int
}

type DoubleZero struct {
  Person
  First string
  LicenseToKill bool
}

func main() {
  p1 := Person{"John", "Doe", 42}
  fmt.Println(p1.First, p1.Last, p1.Age)

  p007 := DoubleZero {
    Person: Person{
      First: "James",
      Last: "Bond",
      Age: 20,
    },
    First: "Double Zero Seven",
    LicenseToKill: true,
  }

  fmt.Println(p007.Person.First, p007.First)
}