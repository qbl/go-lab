package main

import ("fmt"; "math")

type Circle struct {
  x, y, r float64
}

func circleArea(c Circle) float64 {
  return math.Pi * c.r * c.r
}

func setRadius(c *Circle, r float64) {
  c.r = r
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

type Rectangle struct {
  x1, y1, x2, y2 float64
}

func distance(x1, x2, y1, y2 float64) float64 {
  a := x2 - x1
  b := y2 - y1

  return math.Sqrt(a * a + b * b)
}

func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)

  return l * w
}

// Random
type ArrayedStruct struct {
  i [6]int
}

func main() {
  var c1 Circle
  fmt.Println("Circle c1: ", c1)

  // Pointer
  c2 := new(Circle)
  fmt.Println("Circle c2: ", c2)

  c3 := Circle{x: 0, y: 0, r: 5}
  fmt.Println("Circle c3: ", c3)

  c4 := Circle{0, 0, 10}
  fmt.Println("Circle c4: ", c4)

  // Pointer
  c5 := &Circle{0, 0, 15}
  fmt.Println("Circle c5: ", c5)

  fmt.Println("=================")

  fmt.Println(c1.x, c1.y, c1.r)
  c1.x = 1
  c1.y = 2
  fmt.Println(c1.x, c1.y, c1.r)

  fmt.Println("Circle area of c3 is: ", circleArea(c3))

  setRadius(&c4, 20)
  fmt.Println("Circle c4: ", c4)

  fmt.Println("Circle area of c4: ", c4.area())

  r1 := Rectangle{0, 0, 10, 10}
  fmt.Println("Rectangle area of r1: ", r1.area())

  fmt.Println("=================")

  // Arrayed Struct
  as1 := new(ArrayedStruct)
  as1.i = [6]int{1, 2, 3, 4, 5, 6}
  fmt.Println(as1)

  as2 := ArrayedStruct{
    i: [6]int{1, 2, 3, 4, 5, 6},
  }
  fmt.Println(as2)
}
