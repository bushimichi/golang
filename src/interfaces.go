package main

import (
  "fmt"
  "math"
)
  
type I interface {
  M()  // func M 
}

type T struct {
  S string
}

func (t T) M() {
  fmt.Println(t.S)
}

type F float64
func (f F) M() {
  fmt.Println(f)
}

func main() {
  var i I = T{"hello"}
  i.M()
  
  i = F(math.Pi)
  i.M()
  
}