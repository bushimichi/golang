package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

type table struct {
  X, Y int
}

var( 
  v1 = table{1 ,2}
  v2 = table{X: 1}
  v3 = table{}
  p1 = &table{1, 2}
)

func main() {
  fmt.Println(Vertex{1, 2})
  
  v := Vertex{1,2}
  v.X = 4
  
  fmt.Println(v.X)
  
  p := &v
  p.X = 1e9
  fmt.Println(v)
  
  fmt.Println(v1, p1, v2, v3)
}

