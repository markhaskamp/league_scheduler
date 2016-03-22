package main

import (
  "flag"
  "scheduler"
)


func main() {
  n := flag.Int("n", 20, "number of teams")
  flag.Parse()

  scheduler.BuildSchedule(*n)
}


