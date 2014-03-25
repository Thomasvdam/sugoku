package main

import (
 "fmt"
)

/*****
 * A grid consists out of 9 rows, collumns, and clusters.
 *****/
type grid struct {
  rows []*collection
  collumns []*collection
  clusters []*collection
}

/*****
 * Create a new grid with 9 collumns, rows, and clusters based on the string passed.
 *****/
func NewGrid(boardString string) (*grid) {
  grid := &grid{make([]*collection, 9), make([]*collection, 9), make([]*collection, 9)}
  for i := 0; i < 9; i++ {
    grid.rows[i]     = &collection{make([]*square, 9), make(map[int]bool), make(map[int]bool)}
    grid.collumns[i] = &collection{make([]*square, 9), make(map[int]bool), make(map[int]bool)}
    grid.clusters[i] = &collection{make([]*square, 9), make(map[int]bool), make(map[int]bool)}
  }

  pos := 0
  for indexRow := 0; indexRow < 9; indexRow++ {
    for indexCollumn := 0; indexCollumn < 9; indexCollumn++ {
      temp := NewSquare(pos, int(boardString[pos]) - 48)

      // Determine the row/collumn/cluster based on the position in the string.
      rowPos := pos / 9
      collumnPos := pos % 9
      clusterPos := (collumnPos / 3) + (3 * (rowPos / 3))

      // Store the row/collumn/cluster data in the square.
      temp.row = grid.rows[rowPos]
      temp.collumn = grid.collumns[collumnPos]
      temp.cluster = grid.clusters[clusterPos]

      // Store the square in the grid.
      grid.rows[rowPos].squares[collumnPos] = temp
      grid.collumns[collumnPos].squares[rowPos] = temp
      grid.clusters[clusterPos].squares[(rowPos % 3) + (3 * (collumnPos % 3))] = temp

      // Move to the next char in the string.
      pos++
    }
  }

  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      grid.rows[j].UpdateUsed()
      grid.collumns[j].UpdateUsed()
      grid.clusters[j].UpdateUsed()
    }
  }

  return grid
}

/*****
 * Print a readable version of the board.
 *****/
func (grid *grid) Print() {
  for indexRow, row := range grid.rows {
    for indexCollumn, square := range row.squares {
      fmt.Print(square.final, " ")
      if (indexCollumn % 3) == 2 && indexCollumn != 8 {
        fmt.Print("| ")
      }
    }
    fmt.Println()
    if (indexRow % 3) == 2 && indexRow != 8 {
      fmt.Println("----------------------")
    }
  }
  fmt.Println()
}

/*****
 * Runs one update round for all the rows, clusters, and collumns.
 *****/
func (grid *grid) Update() {
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      grid.rows[j].Update()
      grid.collumns[j].Update()
      grid.clusters[j].Update()
    }
  }
}

/*****
 * Check whether the board has been completed.
 *****/
func (grid *grid) Completed() bool {
  for _, row := range grid.rows {
    for _, square := range row.squares {
      if square.final == 0 {
        return false
      }
    }
  }
  return true
}