package main

import (
  "fmt"
)

func Sqrt(x float64) float64 {
  tmp := x
  x1 := x
  for i := 0; i < 10; i++ {
    x1 = x1 - (x1*x1 - x)/( 2 * x1)
    fmt.Println(x1)
    
    if x1 == tmp{
      break
    }
    tmp = x1
  }
  return x1
}

func main() {
  fmt.Println(Sqrt(2))
}