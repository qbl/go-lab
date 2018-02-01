package main

import "fmt"

func swap(x, y *int) {
  temp := new(int)
  *temp = *x
  *x = *y
  *y = *temp
}

func main() {
  x := 1
  y := 2
  fmt.Printf("The value of x is: %d and the value of y is: %d\n", x, y)
  swap(&x, &y)
  fmt.Printf("The value of x is: %d and the value of y is: %d\n", x, y)
}