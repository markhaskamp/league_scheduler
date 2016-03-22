package main

import (
  "flag"
  "scheduler"
)


func main() {
  n := flag.Int("n", 20, "number of teams")

  scheduler.BuildSchedule(*n)
}


