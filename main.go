package main

import (
  "fmt"
  _ "time"
)

func main() {
  r := New()
  // ce := cell{10, make(chan int)}
  ce := r.CreateInput(10)

  compCell := r.CreateCompute1(ce, func (val int) int {
    return val * val
  })

  // fmt.Println(r.CreateInput(10))
  fmt.Println(compCell.Value())
  ce.SetValue(200)
  ce.SetValue(22)
  ce.SetValue(0)
  fmt.Println(compCell.Value())
  // fmt.Println(compCell.Value())
  // time.Sleep(time.Millisecond * 1)
}