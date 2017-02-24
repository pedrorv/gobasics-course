package main

import (
  "encoding/json"
  "strings"
  "fmt"
)

type Person struct {
  First       string
  Last        string
  Age         int
  notExported int
}


func main() {
  var p1 Person
  rdr := strings.NewReader(`{"First":"John","Last":"Doe","Age":42}`)
  json.NewDecoder(rdr).Decode(&p1)

  fmt.Println(p1.First)
  fmt.Println(p1.Last)
  fmt.Println(p1.Age)
}