package main

import "fmt"

func main() {
  value, exists := power("Goku")
  if exists == false {
    fmt.Printf("false")
  }
}

func log(message string) {
  return "hello"
}

func add(a int, b int) int {
  return (a + b)
}

func power(name string) (int, bool) {
  return 9000, true
}