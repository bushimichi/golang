package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

// 同じ型の引数の場合は省略できる
func add2(x,y int) int{
  return x + y
}

// 複数の戻り値を返す場合
func swap(x, y string) (string, string){
  return y, x
}

// Named return values
func split(sum int) (x, y int){
  x = sum * 4 / 9
  y = sum - x
  return
}

func main(){
  fmt.Println("[add] 42 + 13 =", add(42,13))
  fmt.Println("[add2] 42 + 13 =", add2(42,13))

  a, b := swap("hello", "world")
  fmt.Println(a,b)

  fmt.Println(split(17))
}
