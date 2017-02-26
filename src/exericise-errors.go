package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
  return fmt.Sprintf("cannot Sqrt nagative number: %s", float64(e))
}

func Sqrt(x float64) (float64, error) {
  
  if x < 0 {
    return x, ErrNegativeSqrt(x)
  }
  
  tmp := x
  x1 := x
  for i := 0; i < 10; i++ {
    x1 = x1 - (x1*x1 - x)/( 2 * x1)
    //fmt.Println(x1)
    
    if x1 == tmp{
      break
    }
    tmp = x1
  }
  
  return tmp, nil
}

func main(){
  fmt.Println(Sqrt(2))
  fmt.Println(Sqrt(-2))
}