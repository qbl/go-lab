package main

import "fmt"

func main() {
  a := 2
  b := 3
  fmt.Println(add(a, b))
}

func add(args ...int) int {
  total := 0
  for _, v := range(args) {
    total += v 
  }

  return total
}