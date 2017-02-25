package main

import "fmt"

var c, python, java bool
var j, k int = 1, 2

func main(){
  var i int
  fmt.Println(i, c, python, java)

  var c, python, java = true, false, "no!"
  fmt.Println(j, k, c, python, java)
}