package main

import (
  "fmt"
  "runtime"
  "time"
)

func main() {
  fmt.Print("Go runs on ")

  os := runtime.GOOS;
  switch os{
    case "darwin":
      fmt.Println("OS X.")
    case "linux":
      fmt.Println("Linux.")
    default:
    fmt.Printf("%s.", os)
  }
  
  t := time.Now()
  switch{
    
    case t.Hour() < 12:
      fmt.Println("good morning")
    case t.Hour() < 17:
      fmt.Println("good afternoon")
    default:
      fmt.Println("good evening")
  }
 
  t1 := time.Now().Hour()
  switch{
    case t1 < 12:
    fmt.Println("おはようございます")
    case t1 < 17:
    fmt.Println("こんにちは")
    default:
    fmt.Println("こんばんは")
  }
}
