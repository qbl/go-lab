package main

import "fmt"

func half(num int) (int, bool) {
  var isEven bool
  if num % 2 == 0 {
    isEven = true
  } else {
    isEven = false
  }

  return num/2, isEven
}

func greatest(args ...int) int {
  greatest := args[0]
  for _, v := range(args) {
    if v > greatest {
      greatest = v
    }
  }

  return greatest
}

func fibonacci(num int) int {
  val := 0
  if num == 0 {
    val = 0
  } else if num == 1 {
    val = 1
  } else {
    val = fibonacci(num-1) + fibonacci(num-2)
  }

  return val
}

func square(x *float64) {
  *x = *x * *x
}

func swap(x,y *int) {
  temp := new(int)
  *temp = *x
  *x = *y
  *y = *temp
}

func main() {
  fmt.Println(half(1))
  fmt.Println(half(2))

  fmt.Println(greatest(1,5,2,6,3,7,4))

  fmt.Println(fibonacci(3))

  x := 1.5
  square(&x)
  fmt.Println(x)

  a := 1
  b := 2
  fmt.Println(a)
  fmt.Println(b)
  swap(&a, &b)
  fmt.Println(a)
  fmt.Println(b)
}