package main

import "fmt"


func main(){
  s := []int{2,3,5,7,11,13}
  printSlice(s)
  
  // Slice the slice to give it zero length
  s = s[:0]
  printSlice(s)
  
  // Extend its length.
  s = s[:4]
  printSlice(s)
  
  // Drop its first two values.
  s = s[2:]
  printSlice(s)
  
  s = s[0:4]
  printSlice(s)
  
  // zero
  var s1 []int
  fmt.Println(s1, len(s1), cap(s1))
  
  if s1 == nil {
    fmt.Println("nil!")
  }

  a := make([]int, 5)
  printSlice2("a", a)
  
  b := make([]int, 0, 5)
  printSlice2("b", b)
  
  c := b[:2]
  printSlice2("c", c)
  
  d := c[2:5]
  printSlice2("d", d)
  
  
  
}

func printSlice(s []int){
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice2(s1 string, s []int){
  fmt.Printf("%s len=%d cap=%d %v\n",s1, len(s), cap(s), s)
}
