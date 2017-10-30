package main

import "fmt"

func increment(num int) int {
  return num + 1
}

func main() {
  fmt.Printf("My number starts from %d\n", myNumber)
  myNumber = increment(myNumber)
  fmt.Printf("My number is now %d\n", myNumber)
}
