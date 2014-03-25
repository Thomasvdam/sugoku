package main

import (
  "os"
  "fmt"
  "strings"
)

func main() {
  // Get the board from the command line and check it.
  if len(os.Args) != 2 {
    fmt.Println("Please enter an 81 character string, using 0 as whitespace and working from left to right, top to bottom.")
    os.Exit(1)
  }

  boardString := os.Args[1]
  boardReader := strings.NewReader(boardString)
  if boardReader.Len() != 81 {
    fmt.Println("Please enter an 81 character string, using 0 as whitespace and working from left to right, top to bottom.")
    os.Exit(1)
  }

  // Create the grid and print it once for easy error checking.
  grid := NewGrid(boardString)
  grid.Print()
  
  // Update the grid while it is not yet completed.
  for !grid.Completed() {
    grid.Update()
  }

  // Print the solved grid.
  grid.Print()
}
