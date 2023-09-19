package main

import (
  "fmt"
  "os"
)

func main() {
  grid := newMineSweeper()
  grid.Print()
  var buf string
  var bufX, bufY int
  var err error
  var isValid bool
  for {
    fmt.Scan(&buf)
    fmt.Scan(&bufX)
    fmt.Scan(&bufY)
    switch buf {
    case "open":
      isValid, err = grid.OpenCell(bufX - 1, bufY - 1)
      if !isValid {
        os.Exit(1)
      }
    case "flag":
      err = grid.SetFlag(bufX - 1, bufY - 1)
    }
    if err != nil {
      fmt.Println(err)
      break
    }
    if grid.victoryCheck() {
      fmt.Println("Победа!!!")
    }
    grid.Print()
  }
}
