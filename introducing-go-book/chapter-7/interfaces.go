package main

import ("fmt"; "math")

type Shape interface {
  area() float64
}

// Circle
type Circle struct {
  x, y, r float64
}

func (c *Circle) area() float64 {
  return math.Pi * c.r * c.r
}

// Rectangle
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

// Total Area
func totalArea(shapes ...Shape) float64 {
  var area float64
  for _, s := range(shapes) {
    area += s.area()
  }

  return area
}

// Multishapes
type MultiShape struct {
  shapes []Shape
}

func (m *MultiShape) area() float64 {
  var area float64
  for _, v := range(m.shapes) {
    area += v.area()
  }

  return area
}

func main() {
  c := Circle{0, 0, 14}
  r := Rectangle{0, 0, 10, 10}
  fmt.Println(totalArea(&c, &r))

  m := MultiShape{
    shapes: []Shape{
      &Circle{0, 0, 14},
      &Rectangle{0, 0, 10, 10},
    },
  }
  fmt.Println(m)
  fmt.Println(m.area())
}