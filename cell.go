package main

import (
  "errors"
)

var (
  ErrIncorectNumber = errors.New("Неверное число, число должно быть в пределах от -1 до 8")
)


type cell struct {
  val int // -1 - bomb, 0 - 8 - numOfBombs
  isOpen bool
  flag bool
}

func (c *cell) Val() int {
  return c.val
}

func (c *cell) Incr() {
  c.val++
}

func (c *cell) SetVal(num int) error {
  if !isCorrect(num) {
    return ErrIncorectNumber
  }
  c.val = num
  return nil
}

func (c *cell) IsOpen() bool {
  return c.isOpen
}

func (c *cell) SetIsOpen(fl bool) {
  c.isOpen = fl
}

func (c *cell) Flag() bool {
  return c.flag
}

func (c *cell) SetFlag(fl bool) {
  c.flag = fl
}

func isCorrect(num int) bool {
  if num < -1 || num > 8 {
    return false
  }
  return true
}
