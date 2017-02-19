package main

import "fmt" 

func main() {

  // Like a C for
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  j := 0
  // Like a C while 
  for j < 10 {
    fmt.Println(j)
    j++
  }

  k := 0
  // No condition
  for {
    if k >= 20 {
      break
    }

    fmt.Println(k) 
    k++
  }

}