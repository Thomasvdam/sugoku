package main

/*****
 * A square has a list of possible numbers, and pointers to the
 * row, collumn, and cluster it is located in.
 *****/
type square struct {
  pos int
  final int
  numbers map[int]bool

  row *collection
  collumn *collection
  cluster *collection
}

/*****
 * Allocate and initialise a square.
 *****/
func NewSquare(pos, val int) (*square) {
  s := &square{pos, 0, make(map[int]bool), nil, nil, nil}
  
  if val == 0 {
    for i := 1; i < 10; i++ {
      s.numbers[i] = true
    }
  } else {
    s.final = val
  }

  return s
}

/*****
 * Update the possible numbers for the square.
 *****/
func (square *square) Remove(x int) bool {
  if square.final != 0 {
    return false
  }

  delete(square.numbers, x)
  if len(square.numbers) == 1 {
    for key, _ := range square.numbers {
      square.final = key
      return true
    }
  }

  return false
}

/*****
 * Call all relevant UpdateUsed methods.
 *****/
func (self *square) UpdateUsed() {
  self.row.UpdateUsed()
  self.collumn.UpdateUsed()
  self.cluster.UpdateUsed()
}