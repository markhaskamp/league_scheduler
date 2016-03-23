package scheduler

import (
  "fmt"
  "math/rand"
  "time"
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

  // build schedule for inter conference
    // this is the first 10 weeks of the schedule
  var interLeagueSchedule [][]Matchup
  interLeagueSchedule = buildInterLeagueMatchups(int(n/2), l1, l2)

  // build schedules for intra conferences for each conference
  // merge the two schedules
    // this is the last 9 weeks of the schedule
  intraGoldSchedule := buildIntraLeagueMatchups(l1)

  intraSilverSchedule := buildIntraLeagueMatchups(l2)

  // build entire intraLeague schedule
  intraLeagueSchedule := make([][]Matchup, len(intraSilverSchedule))
  for i:=0; i<len(intraSilverSchedule); i++ {
    intraLeagueSchedule[i] = append(intraGoldSchedule[i], intraSilverSchedule[i]...)
    intraLeagueSchedule[i] = randomizeStartTimes(intraLeagueSchedule[i])
  }

  // append intraLeague onto (and after) interLeague
  completeSchedule := append(interLeagueSchedule, intraLeagueSchedule...)
  printWeeklyMatchups(completeSchedule)

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

    weekly = randomizeStartTimes(weekly)

    foo[i] = weekly
  }

  return foo
}

func randomizeStartTimes(matchups []Matchup) []Matchup {
  returnMatchups := make([]Matchup, len(matchups))
  time.Sleep(1234 * time.Millisecond)
  randoms := randomRange(len(matchups))

  for ndx,v := range randoms {
    returnMatchups[ndx] = matchups[v]
  }

  return returnMatchups
}


func buildIntraLeagueMatchups(l1 []int) [][]Matchup {
  // l1 1,2,3,4, 5
  //    6,7,8,9,10

  n := len(l1)
  topRow := l1[0:n/2]
  bottomRow := l1[n/2:n]

  returnArray := make([][]Matchup, 9)

  for i:=0; i<n-1; i++ {

    topIntra := make([]Matchup, n/2)
    for ndx,_ := range(topRow) {
      topIntra[ndx] = Matchup{t1:topRow[ndx], t2:bottomRow[ndx]}
    }
    returnArray[i] = topIntra

    topRow, bottomRow = rotateForIntra(n, topRow, bottomRow)
  }

  return returnArray
}


func rotateForIntra(n int, a []int, b []int) ([]int, []int) {

  var foo int
  foo = n/2 - 1
  
  tmp := a[foo]
  for i:=foo; i>0; i-- {
    a[i] = a[i-1]
  }

  tmp2 := b[0]
  for j:=0; j<foo; j++ {
    b[j] = b[j+1]
  }

  b[foo] = tmp
  a[1] = tmp2

  return a,b
}

func printWeeklyMatchups(matchups [][]Matchup) {
  for _,week := range matchups {
    fmt.Println(week)
  }
}


func randomRange(n int) []int {
  returnArray := make([]int, 0)
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)

  foo := map[int]int{0:0, 1:0, 2:0, 3:0, 4:0, 5:0, 6:0, 7:0, 8:0, 9:0}

  for not10(foo) {
    num := r1.Intn(n)
    if foo[num] == 0 {
      returnArray = append(returnArray, num)
    }
      
    foo[num] = 1
  }

  return returnArray
}


func not10(m map[int]int) bool {
  total := 0

  for _,v := range m {
    total += v
  }

  return total < 10
}
