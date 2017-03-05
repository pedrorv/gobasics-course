package main

import "fmt"

func main() {

  if true {
    fmt.Println("Will print")
  }

  a := true

  // This initialization is scoped only inside the if block
  if b := "Will print"; a {
    fmt.Println(b)
  }

  if false {
    fmt.Println("Won't print")
  } else if true {
    fmt.Println("Will print")
  } else {
    fmt.Println("Won't print")
  }

}