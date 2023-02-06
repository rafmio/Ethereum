package main

import (
  "fmt"
)

type Person struct {
  Name string
  Age int
}

func main() {
  var persons = []Person{
    {"Kenny", 20},
    // {"Menny", 23},
    // {"Johnny", 30},
    // {"Tonny", 29},
    // {"Sindey", 30},
    // {"Mike", 34},
  }

  var indexForDelete int

  for i, value := range persons {
    if value.Name == "Kenny" {
      indexForDelete = i
    }
  }

  fmt.Println("indexForDelete: ", indexForDelete)

  persons = append(persons[:indexForDelete], persons[indexForDelete + 1 :] ...)
  fmt.Println(persons)
}
