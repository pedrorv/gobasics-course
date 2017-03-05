package main

import (
  "encoding/json"
  "fmt"
)

type Person struct {
  First       string
  Last        string
  Age         int
  notExported int
}

type PersonTags struct {
  First string `json:"first"`
  Last string `json:"last"`
  Age int `json:"age"`
}

func main() {
  p1 := Person{"John", "Doe", 42, 84}
  bs, _ := json.Marshal(p1)
  fmt.Println(bs)
  fmt.Printf("%T \n", bs)
  fmt.Println(string(bs))

  p2 := PersonTags{"John", "Doe", 42}
  bs2, _ := json.Marshal(p2)
  fmt.Println(string(bs2))
}