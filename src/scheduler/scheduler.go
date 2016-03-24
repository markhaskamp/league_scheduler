package scheduler

import (
  "fmt"
  "crypto/rand"
  "math/big"
  "strconv"
)


type Matchup struct {
  t1 int
  t2 int
}


func BuildSchedule(n int) {
  if n%2 == 1 {
    n++
  }

  l1, l2 := createTwoLeagues(n)

  var interLeagueSchedule [][]Matchup
  interLeagueSchedule = buildInterLeagueMatchups(int(n/2), l1, l2)

  intraGoldSchedule := buildIntraLeagueMatchups(l1)
  intraSilverSchedule := buildIntraLeagueMatchups(l2)


  intraLeagueSchedule := make([][]Matchup, len(intraSilverSchedule))
  for i:=0; i<len(intraSilverSchedule); i++ {
    intraLeagueSchedule[i] = append(intraGoldSchedule[i], intraSilverSchedule[i]...)
    intraLeagueSchedule[i] = randomizeStartTimes(intraLeagueSchedule[i])
  }

  intraLeagueSchedule = randomizeWeeks(intraLeagueSchedule)

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
  matchups := make([][]Matchup, weeks)

  for i := 0; i<weeks; i++ {
    weekly := make([]Matchup, weeks)

    for j := 0; j<weeks; j++ {
      weekly[j] = Matchup{t1: l1[j], t2: l2[j]}
    }
    tmp := l2[0]
    for j := 0; j<weeks-1; j++ {
      l2[j] = l2[j+1]
    }
    lastSpot := len(l1) - 1
    l2[lastSpot] = tmp

    weekly = randomizeStartTimes(weekly)

    matchups[i] = weekly
  }

  return matchups
}

func randomizeStartTimes(matchups []Matchup) []Matchup {
  returnMatchups := make([]Matchup, len(matchups))
  randoms := randomRange(len(matchups))

  for ndx,v := range randoms {
    returnMatchups[ndx] = matchups[v]
  }

  return returnMatchups
}

func randomizeWeeks(weeks [][]Matchup) [][]Matchup {
  returnWeeks := make([][]Matchup, len(weeks))
  randoms := randomRange(len(weeks))
  fmt.Println("randoms:", randoms)

  for ndx,v := range randoms {
    returnWeeks[ndx] = weeks[v]
  }

  return returnWeeks
}


func buildIntraLeagueMatchups(l1 []int) [][]Matchup {
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

  randomMap := make(map[int]int)
  for i:=0; i<n; i++ {
    randomMap[i] = 0
  }

  for notDone(n, randomMap) {
    num := GetRandomNumber(n)

    if randomMap[num] == 0 {
      returnArray = append(returnArray, num)
    }
      
    randomMap[num] = 1
  }

  return returnArray
}


func notDone(n int, m map[int]int) bool {
  total := 0

  for _,v := range m {
    total += v
  }

  return total < n
}


func GetRandomNumber(n int) int {

    nBig, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
    if err != nil {
      panic(err)
    }

    s := fmt.Sprintf("%v", nBig.Int64())

    returnNum, err := strconv.Atoi(s)
    if err != nil {
      panic(err)
    }

    return returnNum
}

