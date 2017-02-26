package main

import (
  "fmt"
  "math"
)

type Vertex struct{
  X, Y float64
}

func (self Vertex) Abs() float64 {
  return math.Sqrt(self.X * self.X + self.Y * self.Y)
}

func Abs(self Vertex) float64 {
  return math.Sqrt(self.X * self.X + self.Y * self.Y)
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}


func main(){
  v := Vertex{ 3, 4 }
  fmt.Println(v.Abs())
  fmt.Println(Abs(v))

  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())
}
