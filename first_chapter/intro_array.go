package main

import "fmt"

func main() {
  var x [5]int // Defaults to 0 for all values
  x[4] = 100
  fmt.Println(x)
}