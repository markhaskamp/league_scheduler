package main

import (
  "flag"
  "fmt"
  "scheduler"
)


func main() {
  n := flag.Int("n", 20, "number of teams")
  flag.Parse()
  fmt.Println("got to here")
  fmt.Println("n:", *n)

  scheduler.BuildSchedule(*n)
}


