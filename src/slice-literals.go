package main

import "fmt"

func main(){
  
  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)
  
  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)
  
  s := []struct{
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(s)

  s1 := []int{2, 3, 5, 7, 11, 13}
  
  s1 = s1[1:4]
  fmt.Println(s1)
  
  s1 = s1[:2]
  fmt.Println(s1)
  
  s1 = s1[1:2]
  fmt.Println(s1)
  
  

}