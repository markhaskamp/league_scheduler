package scheduler

import (
  "fmt"
)


func BuildSchedule(n int) {

  // split into two conferences
  c1, c2 := createTwoConferences(n)
  fmt.Println(c1,c2)

  // build schedule for inter conference
    // this is the first 10 weeks of the schedule
  // build schedules for intra conferences for each conference
  // merge the two schedules
    // this is the last 9 weeks of the schedule
}

func createTwoConferences(n int) ([]int, []int) {
  if n %2 == 1 {
    n++
  }

  var x int
  x = n/2;

  c1 := make([]int, x)
  for i := 0; i<x; i++ {
    c1[i] = i + 1
  }

  c2 := make([]int, x)
  for i := x+1; i<=n; i++ {
    c2[i-x-1] = i
  }

  return c1, c2
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
