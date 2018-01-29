package main

import "fmt"

func main() {
  fmt.Println("Enter a number in inch:")

  var input float64
  fmt.Scanf("%f", &input)

  meter := input * 0.3048
  fmt.Println(meter)
}