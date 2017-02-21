package main

import "fmt" 

func main() {

  slice1 := []int{1, 2, 3, 4, 5}
  fmt.Printf("%T\n", slice1)
  fmt.Println(slice1)

  fmt.Println(slice1[1:3])
  fmt.Println(slice1[1])
  fmt.Println("string"[1])

  mySlice := make([]int, 0, 2)

  fmt.Println("---------------")
  fmt.Println(mySlice)
  fmt.Println(len(mySlice))
  fmt.Println(cap(mySlice))
  fmt.Println("---------------")

  for i := 0; i < 80; i++ {
    mySlice = append(mySlice, i)
    fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), 
                "Value:", mySlice[i])
  }

  slice2 := []int{1, 2, 3}
  slice3 := []int{4, 5, 6}

  slice2 = append(slice2, slice3...)

  fmt.Println(slice2)

  // deleting from a slice

  slice4 := []string{"Monday", "Tuesday"}
  slice5 := []string{"Wednesday", "Thursday", "Friday"}

  slice6 := append(slice4, slice5...)
  slice6 = append(slice6[:2], slice6[3:]...)
  fmt.Println(slice6)

  // a slice of string slices
  // slice := make([][]string, 0)
  // slice = append(slice, []string{"some slice"})

  // Creating a slice
  // slice := []string{}
  // var slice []string
  // slice := make([]string, 0) 
}