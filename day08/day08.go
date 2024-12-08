package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Pos [2]int

func main() {
	f, _ := os.Open("./day08/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	mapp, antPos := parseMap(sc)

	cnt := 0
	for _, v := range antPos {
		for i, f := range v {
			for j, ff := range v {
				if i == j {
					continue
				}
				dx := f[0] - ff[0]
				dy := f[1] - ff[1]

				if (f[1]+dy >= 0 && f[1]+dy < len(mapp)) && (f[0]+dx >= 0 && f[0]+dx < len(mapp[0])) && mapp[f[1]+dy][f[0]+dx] != '#' {
					mapp[f[1]+dy][f[0]+dx] = '#'
					cnt++
				}
			}
		}
	}
	return cnt
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	mapp, antPos := parseMap(sc)

	pts := map[Pos]bool{}
	for _, v := range antPos {
		for i, f := range v {
			pts[f] = true // this make it work some how
			for j, ff := range v {
				if i == j {
					continue
				}
				dx := f[0] - ff[0]
				dy := f[1] - ff[1]

				cx := f[0] + dx
				cy := f[1] + dy

				for (cy >= 0 && cy < len(mapp)) && (cx >= 0 && cx < len(mapp[0])) {
					pts[Pos{cx, cy}] = true
					mapp[cy][cx] = '#'
					cx += dx
					cy += dy
				}
			}
		}
	}
	// printGrid(mapp)
	return len(pts)
}

func parseMap(sc *bufio.Scanner) ([][]rune, map[rune][]Pos) {
	antPos := map[rune][]Pos{}
	mapp := make([][]rune, 0)
	for sc.Scan() {
		l := sc.Text()
		mapp = append(mapp, []rune(l))

		for i, m := range l {
			if m == '.' {
				continue
			}

			if _, ok := antPos[m]; !ok {
				antPos[m] = make([]Pos, 0)
			}
			antPos[m] = append(antPos[m], Pos{i, len(mapp) - 1})
		}
	}
	return mapp, antPos
}

func printGrid(g [][]rune) {
	for _, i := range g {
		fmt.Println(string(i))
	}
}
