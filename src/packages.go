package main

import(
  "fmt"
  "math/rand"
  "math"
)

func main(){
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Printf("Now you have %g problems.", math.Sqrt(7))
  
  // 初めの1文字が大文字の場合は外部参照可能
  fmt.Println(math.Pi)
  // 初めの1文字が小文字の場合は外部参照不可能
  // fmt.Println(math.pi)  <= error
  
}
