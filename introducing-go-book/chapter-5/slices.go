package main

import "fmt"

func main() {
  // Slicing basics
  primes := [6]int{2, 3, 5, 7, 11, 13}

  var s []int = primes[1:4]
  fmt.Println(s)

  var s1 []int = primes[1:]
  fmt.Println(s1)

  var s2 []int = primes[:4]
  fmt.Println(s2)

  var s3 []int = primes[:]
  fmt.Println(s3)

  // Append
  slice1 := []int{1, 2, 3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice2)

  // Copy
  slice3 := []int{1, 2, 3}
  slice4 := make([]int, 2)
  copy(slice4, slice3)
  fmt.Println(slice3, slice4)
}