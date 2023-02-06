package main

import "fmt"

func main() {
  numbers := []int{10, 20, 30, 40, 50, 90, 60}
  fmt.Println("Original slise: ", numbers)

  var index int = 3
  fmt.Println("numbers[index] = ", numbers[index])

  numbers = append( numbers[ :(index - 1)], numbers[(index + 1) : ] ...)

  fmt.Println(numbers)
}
