package main

import "fmt"


// Standard function
func zero(x int) {
  x = 0
}

// Function with a pointer
func one(x *int) {
  *x = 1
}

func two(x *int) {
  *x = 2
}

func main() {
  x := 5
  zero(x)
  fmt.Println(x)

  one(&x)
  fmt.Println(x)

  xPtr := new(int)
  fmt.Println(*xPtr)
  two(xPtr)
  fmt.Println(*xPtr)
}