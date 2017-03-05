package main

import "fmt" 

func main() {
  x := 1

  // No need to add break statements 
  // Fallthrough is optional. Needs to be specified
  // with a fallthrough statement

  switch x {
    case 0:
      fmt.Println(0)
    case 1:
      fmt.Println(1)
    case 2:
      fmt.Println(2)
    
    // Multiple cases
    case 3, 4, 5:
      fmt.Println("bigger than 2")
      fallthrough
    case 6, 7, 8, 9:
      fmt.Println("also bigger than 2")
    default:
      fmt.Println("default")
  }

  name := "john"

  switch {
    case name == "john":
      fmt.Println(name)
      fallthrough
    case len(name) == 4:
      fmt.Println("length: ", len(name))
      fallthrough
    default:
      fmt.Println("default")
  }

  value := 42

  

  typeSwitch(value)
}

func typeSwitch(x interface{}) {
  switch x.(type) {
    case int:
      fmt.Println("int")
    case float64:
      fmt.Println("float64")
    default:
      fmt.Println("default")
  }
}