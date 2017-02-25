package main

import (
  "fmt"
  "math/cmplx"
  "math"
)

var (
  ToBe bool = false
  MaxInt uint64 = 1<<64 - 1
  z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
  const f = "%T(%v)\n"
  fmt.Printf(f, ToBe, ToBe)
  fmt.Printf(f, MaxInt, MaxInt)
  fmt.Printf(f, z, z)
  
  var i int
  var fl float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, fl, b, s)
  
  var x, y int = 3,4
  var l float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(l)
  fmt.Println(x, y, z)
  
  x2 := float64(x)
  fmt.Printf("%T(%v)\n", x2, x2)
}