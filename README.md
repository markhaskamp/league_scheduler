# Schedule Builder for Wednesday Golf League

Currently only works for leagues whose member count is evenly divisible by 4

## tech stack

- go

## Notes

- 2 divisions
- 1st half of year is inter division play
- 2nd half of year is intra division play

- intra league algorithm
  - top row is    1,2,3,4,5
  - bottom row is 6,7,8,9,10
  - '1' is anchored
  - 2-5,10-6 is rotated clockwise 9 times to get 9 weeks of different matchups
- inter league algorithm
  - top row is     1-10
  - bottom row is 11-20
  - top row slides to the right 10 times to get 10 weeks of different matchups


