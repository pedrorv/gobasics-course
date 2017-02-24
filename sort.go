package main

import (
  "fmt"
  "sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }



func main() {
  p := people{"Jack", "John", "Jimmy", "Al"}

  fmt.Println(p)
  sort.Sort(p)
  fmt.Println(p)

  s := []string{"Jack", "John", "Jimmy", "Al"}

  fmt.Println(s)
  sort.StringSlice(s).Sort()
  fmt.Println(s)

  s2 := []string{"Jack", "John", "Jimmy", "Al"}

  fmt.Println(s2)
  sort.Strings(s2)
  fmt.Println(s2)

  sort.Sort(sort.Reverse(sort.StringSlice(s2)))
  fmt.Println(s2)

  n := []int{6, 4, 2, 1, 7, 7, 9, 10, 3}

  fmt.Println(n)
  sort.Ints(n)
  fmt.Println(n)
  sort.Sort(sort.Reverse(sort.IntSlice(n)))
  fmt.Println(n)
}