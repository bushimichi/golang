package main

import (
  "fmt"
  "runtime"
  "time"
)

func main() {
  fmt.Print("Go runs on ")
  
  switch os := runtime.GOOS; os {
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
  
}
  
