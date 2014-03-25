package main

/*****
 * A collection has pointers to 9 squares and a map of the numbers already present.
 *****/
type collection struct {
  squares []*square
  used map[int]bool
  unused map[int]bool
}

/*****
 * Update the map containing the used numbers.
 *****/
func (self *collection) UpdateUsed() {
  for _, square := range self.squares {
    if square.final != 0 {
      self.used[square.final] = true
    }
  }

  for i := 1; i < 10; i++ {
    _, ok := self.used[i]
    if ok {
      delete(self.unused, i)
    } else {
      self.unused[i] = true
    }
  }
}

/*****
 * Update all the squares and the used map.
 *****/
func (self *collection) Update() {
  for _, square := range self.squares {
    if square.final == 0 {
      for key, _ := range self.used {
        if square.Remove(key) {
          square.UpdateUsed()
        }
      }
    }
  }

  var possibilities int
  var temp *square
  for key, _ := range self.unused {
    possibilities = 0
    temp = nil
    for _, square := range self.squares {
      if square.final == 0 {
        _, ok := square.numbers[key]
        if ok {
          possibilities++
          temp = square
        }
      }
    }
    if possibilities == 1 {
      temp.final = key
      temp.UpdateUsed()
    }
  }
}
