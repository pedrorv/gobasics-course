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

func (p Person) Greeting() {
  fmt.Println("I'm a regular person")
}

func (dz DoubleZero) Greeting() {
  fmt.Println("I'm an agent")
}

func main() {
  p1 := Person{"John", "Doe", 42}
  
  p007 := DoubleZero {
    Person: Person{
      First: "James",
      Last: "Bond",
      Age: 20,
    },
    First: "Double Zero Seven",
    LicenseToKill: true,
  }

  p1.Greeting()
  p007.Greeting()
  p007.Person.Greeting()

  p2 := &Person{"Jack", "Bauer", 42}
  fmt.Println(p2)
  fmt.Printf("%T\n", p2)
  fmt.Println(p2.First)
}