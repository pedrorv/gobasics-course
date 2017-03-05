package main

import "fmt" 

const (
  _ = iota
  KB = 1 << (iota * 10)
  MB = 1 << (iota * 10)
  GB = 1 << (iota * 10)
)

func main() {
  fmt.Println("binary\t\tdecimal")
  fmt.Printf("%b\t\t%d\n", KB, KB)
  fmt.Printf("%b\t\t%d\n", MB, MB)
  fmt.Printf("%b\t\t%d\n", GB, GB)
}