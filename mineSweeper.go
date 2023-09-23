package main

import (
  "fmt"
  "errors"
  "math/rand"
  "time"
)

const (
  height     =  8
  width      =  8
  bombsCount = 10
)

var (
  ErrOutOfBounds = errors.New("Неверные координаты")
)

const (
  cellIsOpen = "Клетка уже открыта "
  flagOnCell = "Данная клетка помечена флагом "
  gameOver   = "В клетке была бомба, игра закончена "
)

type mineSweeper struct {
  Grid [][]cell
  BombCount int
}

func newMineSweeper() mineSweeper {
  grid := make([][]cell, height)
  for i := 0; i < height; i++ {
    grid[i] = make([]cell, width)
  }
  res := mineSweeper{Grid : grid}
  res.generateBombs()
  return res
}

func (m *mineSweeper) Print() {
  fmt.Printf("\033c")
  fmt.Print("  ")
  for i := 0; i < height; i++ {
    fmt.Printf("%x ", i + 1)
  }
  fmt.Println()
  for i := 0; i < height; i++ {
    fmt.Printf("%x ", i + 1)
    for j := 0; j < width; j++ {
      if m.Grid[i][j].Flag() {
        fmt.Print("p ")
      } else if !m.Grid[i][j].IsOpen() {
        fmt.Print("# ")
      } else if m.Grid[i][j].Val() == -1 {
        fmt.Print("b ")
      }  else {
        fmt.Printf("%d ", m.Grid[i][j].Val())
      }
    }
    fmt.Print("\n")
  }
  fmt.Printf("Number of bombs on grid: %d\n", bombsCount - m.BombCount)
}

func (m *mineSweeper) SetFlag(x, y int) error {
  if !inBounds(x, y) {
    return ErrOutOfBounds
  }
  if m.Grid[x][y].IsOpen(){
    fmt.Println(cellIsOpen)
    return nil
  }
  if m.Grid[x][y].Flag(){
    m.Grid[x][y].SetFlag(false)
    if m.Grid[x][y].Val() == -1 {
      m.BombCount--
    }
  } else {
    m.Grid[x][y].SetFlag(true)
    if m.Grid[x][y].Val() == -1 {
      m.BombCount++
    }
  }
  return nil
}

func (m *mineSweeper) Click(x, y int) (bool, error) {
  if !inBounds(x, y) {
    return true, ErrOutOfBounds
  }
  if !m.Grid[x][y].IsOpen() {
    return m.openCell(x, y), nil
  } else {
    return m.openNeighbors(x, y), nil
  }

}

func (m *mineSweeper) openNeighbors(x, y int) bool{
  neighborBombs := 0
  for dx := -1; dx <= 1; dx++ {
    newX := x + dx
    for dy := -1; dy <= 1; dy++ {
      newY := y + dy
      if !inBounds(newX, newY) {
        continue
      }
      if !m.Grid[newX][newY].IsOpen() && m.Grid[newX][newY].Flag() {
        neighborBombs++
      }
    }
  }
  if neighborBombs != m.Grid[x][y].Val() {
    return true
  }
  for dx := -1; dx <= 1; dx++ {
    newX := x + dx
    for dy := -1; dy <= 1; dy++ {
      newY := y + dy
      if !inBounds(newX, newY) {
        continue
      }

      if !m.openCell(newX, newY) {
        return false
      }
    }
  }
  return true
}

func (m *mineSweeper) openCell(x, y int) bool {
  if m.Grid[x][y].Flag() {
    fmt.Println(flagOnCell)
    return true
  }
  if m.Grid[x][y].Val() == -1 {
    fmt.Println(gameOver)
    return false
  }
  m.Grid[x][y].SetIsOpen(true)
  if m.Grid[x][y].Val() == 0 {
    m.dfs(x, y)
  }
  return true
}

func (m *mineSweeper) dfs(x, y int) error{
  if !inBounds(x, y) {
    return ErrOutOfBounds
  }
  for dx := -1; dx <= 1; dx++ {
    for dy := -1; dy <= 1; dy ++ {
      if m.isCorrect(x + dx, y + dy) {
        m.Grid[x + dx][y + dy].SetIsOpen(true)
        if m.Grid[x + dx][y + dy].Val() ==  0 {
          m.dfs(x + dx, y + dy)
        }
      }
    }
  }
  return nil
}

func (m *mineSweeper) isCorrect(x, y int) bool{
    if !inBounds(x, y) {
      return false
    }
    if m.Grid[x][y].Flag() {
      return false
    }
    if m.Grid[x][y].IsOpen() {
      return false
    }
    return true
}

func (m *mineSweeper) generateBombs() {
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < bombsCount; i++ {
    x := rand.Intn(height)
    y := rand.Intn(width)
    if m.Grid[x][y].Val() == -1 {
      i--
    } else {
      m.Grid[x][y].SetVal(-1)
      m.setVals(x, y)
    }
  }
}

func (m *mineSweeper) setVals(x, y int) {
  for dx := -1; dx <= 1; dx++ {
    for dy := -1; dy <= 1; dy++ {
      if x + dx >= 0 && x + dx < height && y + dy >= 0 && y + dy < width {
        if m.Grid[x + dx][y + dy].Val() >= 0 {
          m.Grid[x + dx][y + dy].Incr()
        }
      }
    }
  }
}

func (m *mineSweeper) victoryCheck() bool{
  return m.BombCount == bombsCount
}

func inBounds(x, y int) bool {
  if x >= height || x < 0 || y >= width || y < 0 {
    return false
  }
  return true
}
