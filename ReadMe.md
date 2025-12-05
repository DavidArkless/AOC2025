# Advent of Code 2025 🎄

Solutions for [Advent of Code 2025](https://adventofcode.com/2025), implemented in Go.

## Requirements

- Go 1.22+ (or any recent Go version)
- A terminal (PowerShell on Windows is fine)

## Project Structure

Each day lives in its own directory under the project root.


New days will be added as `day2/`, `day3/`, … `day25/` following the same pattern.

## How to Run

From the project root, run a specific day with Go’s module-aware `run` command. For example, Day 1:

```powershell
go run ./day1
```

Notes:
- Some days may read their input from a file inside the day’s directory. For Day 1, the program reads `day1/rotations.txt`. Keep your puzzle input in that file before running.
- Run commands from the project root so relative paths resolve correctly.

## Day 1 — Rotations on a Circular Track

Input format (one instruction per line):

```
R<number>
L<number>
```

Where `R` means move right (clockwise) and `L` means move left (counter‑clockwise) by the given number of steps.

Program behaviour (as implemented):
- The track has positions `0` through `99` (100 positions total).
- Start position is `50` for both parts.
- Part 1 moves stepwise and counts how many times you land exactly on position `0` after applying each instruction (with wraparound).
- Part 2 counts how many times you pass through position `0`, including wraparounds that skip over `0`.

Output:
- Two lines are printed:
  1. `part1FinalPosition part1ZeroHits`
  2. `part2FinalPosition part2ZeroPasses`


## Progress

- [x] Day 1
- [x] Day 2
- [x] Day 3
- [x] Day 4
- [ ] Day 5
- [ ] Day 6
- [ ] Day 7
- [ ] Day 8
- [ ] Day 9
- [ ] Day 10
- [ ] Day 11
- [ ] Day 12

