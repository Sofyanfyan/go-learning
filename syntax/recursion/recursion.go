package rucursion

import (
	"fmt"
)

// Factorial using recursion
func factorial_recursion(x int) (y int) {
  if x > 0 {
    y = x * factorial_recursion(x-1)
    } else {
      y = 1
    }
  return
}

func Recursion() {
  fmt.Println(factorial_recursion(5))
}