package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100
  fmt.Println(x)

  var y[5]float64
  y[0] = 21
  y[1] = 32
  y[2] = 52
  y[3] = 14
  y[4] = 61
  
  total := 0.0
  for i := 0; i < len(y); i++ {
    total += y[i]
  }
  fmt.Println(total/float64(len(y)))

  total = 0.0
  for _, value := range(y) {
    total += value
  }
  fmt.Println(total/float64(len(y)))
}