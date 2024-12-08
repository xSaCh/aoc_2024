package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Pos [2]int

func main() {
	f, _ := os.Open("./day06/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	mapp, ePos := parseMap(sc)

	cnt := 0
	for (ePos[1] < len(mapp) && ePos[1] >= 0) && (ePos[0] < len(mapp[0]) && ePos[0] >= 0) {

		// oldPos := ePos
		dir := mapp[ePos[1]][ePos[0]]
		out := false
		dir, ePos, out = traverse(dir, ePos, mapp)
		if out {
			break
		}

		if mapp[ePos[1]][ePos[0]] == '.' {
			cnt++
		}
		// mapp[oldPos[1]][oldPos[0]] = 'X'
		mapp[ePos[1]][ePos[0]] = dir
	}
	return cnt + 1
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	cnt := 0
	mapp, pos := parseMap(sc)

	path := make(map[Pos]bool)
	ePos := pos
	dir := mapp[ePos[1]][ePos[0]]

	for (ePos[1] < len(mapp) && ePos[1] >= 0) && (ePos[0] < len(mapp[0]) && ePos[0] >= 0) {

		out := false
		dir, ePos, out = traverse(dir, ePos, mapp)
		if out {
			break
		}

		path[ePos] = true
	}

	for p := range path {
		cpy := deepClone2D(mapp)
		tmpPath := make(map[Pos]string)
		ePos = pos
		if p == ePos {
			continue
		}
		cpy[p[1]][p[0]] = '#'

		for (ePos[1] < len(cpy) && ePos[1] >= 0) && (ePos[0] < len(cpy[0]) && ePos[0] >= 0) {
			// oldPos := Pos{ePos[0], ePos[1]}
			dir := cpy[ePos[1]][ePos[0]]

			out := false
			dir, ePos, out = traverse(dir, ePos, cpy)
			if out {
				break
			}

			if cDir, exists := tmpPath[ePos]; exists && strings.Contains(cDir, string(dir)) {
				// cpy[p[1]][p[0]] = 'O'
				// printGrid(cpy)
				cnt++
				break
			}

			tmpPath[ePos] = tmpPath[ePos] + string(dir)
			// cpy[oldPos[1]][oldPos[0]] = 'X'
			cpy[ePos[1]][ePos[0]] = dir
		}
	}
	return cnt
}

func traverse(dir rune, ePos Pos, mapp [][]rune) (rune, Pos, bool) {
	switch dir {
	case '^':
		if ePos[1]-1 < 0 {
			return dir, ePos, true
		}
		if mapp[ePos[1]-1][ePos[0]] == '#' {
			dir = '>'
		} else {
			ePos[1] -= 1
		}
	case 'v':
		if ePos[1]+1 >= len(mapp) {
			return dir, ePos, true
		}
		if mapp[ePos[1]+1][ePos[0]] == '#' {
			dir = '<'
		} else {
			ePos[1] += 1
		}
	case '>':
		if ePos[0]+1 >= len(mapp[0]) {
			return dir, ePos, true
		}
		if mapp[ePos[1]][ePos[0]+1] == '#' {
			dir = 'v'
		} else {
			ePos[0] += 1
		}
	case '<':
		if ePos[0]-1 < 0 {
			return dir, ePos, true
		}
		if mapp[ePos[1]][ePos[0]-1] == '#' {
			dir = '^'
		} else {
			ePos[0] -= 1
		}

	}
	return dir, ePos, false
}
func parseMap(sc *bufio.Scanner) ([][]rune, Pos) {
	var enemyPos Pos
	mapp := make([][]rune, 0)
	for sc.Scan() {
		l := sc.Text()
		mapp = append(mapp, []rune(l))

		if idx := strings.IndexFunc(l, func(m rune) bool {
			if m == '^' || m == '>' || m == '<' || m == 'v' {
				return true
			}
			return false
		}); idx != -1 {
			enemyPos[0] = idx
			enemyPos[1] = len(mapp) - 1

		}
	}
	return mapp, enemyPos
}

func printGrid(g [][]rune) {
	for _, i := range g {
		fmt.Println(string(i))
	}
}

func deepClone2D(slice [][]rune) [][]rune {
	clone := make([][]rune, len(slice))
	for i, inner := range slice {
		clone[i] = slices.Clone(inner)
	}
	return clone
}
