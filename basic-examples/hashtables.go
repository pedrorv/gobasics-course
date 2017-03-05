package main

import (
  "bufio"
  "fmt"
  "log"
  "net/http"
)

func main() {
  res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
  if err != nil {
    log.Fatal(err)
  }

  scanner := bufio.NewScanner(res.Body)
  defer res.Body.Close()

  scanner.Split(bufio.ScanWords)
  
  bucketsFirstLetter := make([]int, 200)
  bucketsRemainder := make([]int, 12)
  bucketsEven := make([]int, 12)

  for scanner.Scan() {
    n := HashBucketFirstLetter(scanner.Text())
    m := HashBucketRemainder(scanner.Text(), 12)
    o := HashBucketEven(scanner.Text(), 12)
    bucketsFirstLetter[n]++
    bucketsRemainder[m]++
    bucketsEven[o]++
  }

  fmt.Println(bucketsFirstLetter[65:123])
  fmt.Println(bucketsRemainder)
  fmt.Println(bucketsEven)
}

func HashBucketFirstLetter(word string) int {
  return int(word[0])
}

func HashBucketRemainder(word string, buckets int) int {
  letter := int(word[0])
  bucket := letter % buckets
  return bucket
}

func HashBucketEven(word string, buckets int) int {
  var sum int
  for _, v := range word {
    sum += int(v)
  }
  return sum % buckets
}