package main

import "fmt"

func main() {
  x := 3
  fmt.Println(factorial(x))
}

func factorial(x int) int {
  if x == 0 {
    return 1
  } else {
    return x * factorial(x-1)
  }
}