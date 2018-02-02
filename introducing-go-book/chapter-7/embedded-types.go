package main

import "fmt"

type Person struct {
  Name string
}

func (p *Person) Talk() {
  fmt.Println("Hello my name is", p.Name)
}

type Android struct {
  Person
  Model string
}

func main() {
  person := Person{"Iqbal"}
  person.Talk()

  android1 := new(Android)
  android1.Person.Talk()

  android2 := Android{person, "A123"}
  android2.Talk()
}