package scheduler

import (
  "fmt"
)


type Matchup struct {
  t1 int
  t2 int
}


func BuildSchedule(n int) {
  if n%2 == 1 {
    n++
  }

  // split into two leagues
  l1, l2 := createTwoLeagues(n)
  fmt.Println(l1,l2)

  // build schedule for inter conference
    // this is the first 10 weeks of the schedule
  var interLeagueSchedule [][]Matchup
  interLeagueSchedule = buildInterLeagueMatchups(int(n/2), l1, l2)
  fmt.Println(interLeagueSchedule)

  // build schedules for intra conferences for each conference
  // merge the two schedules
    // this is the last 9 weeks of the schedule
}

func createTwoLeagues(n int) ([]int, []int) {
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



func buildInterLeagueMatchups(weeks int, l1 []int, l2 []int) [][]Matchup {
  foo := make([][]Matchup, weeks)

  for i := 0; i<weeks; i++ {
    weekly := make([]Matchup, weeks)

    for j := 0; j<weeks; j++ {
      weekly[j] = Matchup{t1: l1[j], t2: l2[j]}
    }
    tmp := l2[0]
    for j := 0; j<weeks-1; j++ {
      l2[j] = l2[j+1]
    }
    l2[9] = tmp


    foo[i] = weekly
  }

  return foo
}


