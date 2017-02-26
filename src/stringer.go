package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func (self Person) String() string{
  return fmt.Sprintf("%v (%v years)", self.Name, self.Age)
}

func main() {
  a := Person{"Arthur Dent", 42}
  z := Person{"Zaphod Beeblebrox", 9001}
  
  fmt.Println(a)
  fmt.Println(z)
}