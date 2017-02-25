package main

import "fmt"

func main(){
  var i=1
  for{
    i+=1
    if i > 100 {
      fmt.Println(i)
      return
    }
  }
}
