package main

import "fmt"

func main() {
  addition := func(x, y int) int {
    return x + y
  }
  fmt.Println(addition(10, 20))
}