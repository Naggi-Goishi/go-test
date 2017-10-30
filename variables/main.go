package main

import "fmt"

var (
  name string
  age int
  nationality string
)

func main() {
  name = "Naggi Goishi"
  age  = 19
  nationality = "Japan"

  fmt.Printf("%s (%d) of %s", name, age, nationality)
}
