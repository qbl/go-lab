package main

import "fmt"

func main() {
  x := make(map[string]int)
  
  x["ten"] = 10
  fmt.Println(x)
  fmt.Println(x["ten"])
  fmt.Println(len(x))

  x["nine"] = 9
  fmt.Println(len(x))

  delete(x, "ten")
  fmt.Println(x)

  delete(x, "eight")
  fmt.Println(x)

  fmt.Println(x["seven"])
  fmt.Println(x["ten"])

  if element, ok := x["seven"]; ok {
    fmt.Println(element, ok)
  }

  if element, ok := x["nine"]; ok {
    fmt.Println(element, ok)
  }

  elements := map[string]string {
    "H": "Hydrogen",
    "O": "Oxygen",
    "N": "Nitrogen",
  }
  fmt.Println(elements)

  cheElements := map[string]map[string]string {
    "H" : map[string]string {
      "name": "Hydrogen",
      "state": "Gas",
    },
  }
  fmt.Println(cheElements)

  y := make(map[int]int)
  y[1] = 100
  fmt.Println(y[1])

  z := [6]string {"a", "b", "c", "d", "e", "f"}
  fmt.Println(z[2:5])

  a := []int {
    48, 96, 86, 68,
    57, 82, 63, 70,
    37, 34, 83, 27,
    19, 97, 9, 17,
  }
  min := a[0]
  for _, value := range(a) {
    if value < min {
      min = value
    }
  }
  fmt.Println(min)
}