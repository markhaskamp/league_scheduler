package scheduler

import (
  "fmt"
)


func BuildSchedule(n int) {
  fmt.Println("n:", n)
}


/*
func buildRoundRobinSchedule(teams []int) {
  if len(teams) % 2 != 0 {
    teams[len(teams)] = -1
  }
  n := len(teams)

  // split teams into two slices
  top    := teams[0:n/2]
  bottom := teams[n/2:n]

  schedule := make([]Matchup, 0)
  for ndx,_ range top {
    m := getMatchup(ndx, top, bottom)
  }

  // [0]'s play each other, [1]s play each other, etc
  // slice[0][1] is the anchor
  // rotate all the other array elements
  // split and pair

}
*/
