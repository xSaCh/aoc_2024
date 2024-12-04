package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("./day04/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	data := make([][]rune, 0)
	for sc.Scan() {
		l := sc.Text()
		data = append(data, []rune(l))
	}
	cnt := 0
	for r := 0; r < len(data); r++ {
		for c := 0; c < len(data[0]); c++ {
			if data[r][c] != 'X' {
				continue
			}
			cnt += check(data, r, c, +1, 00) // UP
			cnt += check(data, r, c, -1, 00) // DOWN
			cnt += check(data, r, c, 00, +1) // RIGHT
			cnt += check(data, r, c, 00, -1) // LEFT
			cnt += check(data, r, c, +1, +1) // D RIGHT
			cnt += check(data, r, c, -1, -1) // D LEFT
			cnt += check(data, r, c, +1, -1) // RD RIGHT
			cnt += check(data, r, c, -1, +1) // RD LEFT
		}
	}
	return cnt
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	data := make([][]rune, 0)
	for sc.Scan() {
		l := sc.Text()
		data = append(data, []rune(l))
	}
	cnt := 0
	for r := 0; r < len(data); r++ {
		for c := 0; c < len(data[0]); c++ {
			if data[r][c] != 'A' {
				continue
			}

			// \
			xc := 0
			if check2(data, r-1, c-1, +1, +1) ||
				check2(data, r+1, c+1, -1, -1) {
				xc++
			}

			// /
			if check2(data, r+1, c-1, -1, +1) ||
				check2(data, r-1, c+1, +1, -1) {
				xc++
			}
			if xc == 2 {
				cnt++
			}
		}
	}
	return cnt
}

func check(data [][]rune, r, c, rd, cd int) int {
	if getAt(data, r+1*rd, c+1*cd) == 'M' && getAt(data, r+2*rd, c+2*cd) == 'A' && getAt(data, r+3*rd, c+3*cd) == 'S' {
		return 1
	}
	return 0
}

func check2(data [][]rune, r, c, rd, cd int) bool {
	return getAt(data, r+0*rd, c+0*cd) == 'M' && getAt(data, r+2*rd, c+2*cd) == 'S'
}

func getAt(d [][]rune, r, c int) rune {
	if r < 0 || r >= len(d) {
		return '.'
	}
	if c < 0 || c >= len(d[r]) {
		return '.'
	}

	return d[r][c]
}
