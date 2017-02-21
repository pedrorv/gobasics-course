package main

import "fmt"

func main() {
  m := make(map[string]int)
  m["key1"] = 42
  m["key2"] = 84
  fmt.Println(m)
  
  delete(m, "key2")
  fmt.Println(m)

  m2 := map[string]int{"key1": 42, "key2": 84}
  fmt.Println(m2)


  numbers := map[string]int {
    "zero": 0,
    "one": 1,
    "two": 2,
  }

  fmt.Println(numbers)

  if val, exists := numbers["one"]; exists {
    fmt.Println("val:", val)
    fmt.Println("exists:", exists)
  }

  for key, val := range numbers {
    fmt.Println(key, "-", val)
  }
}