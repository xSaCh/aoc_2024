package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day07/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)

	// oprands := map[int][]int{}

	sum := 0
	for sc.Scan() {
		l := sc.Text()
		sl := strings.Split(l, ": ")

		res, _ := strconv.Atoi(sl[0])
		opds := []int{}
		for _, opd := range strings.Split(sl[1], " ") {
			o, _ := strconv.Atoi(opd)
			opds = append(opds, o)
		}

		possibleAns := []int{opds[0]}
	outer:
		for i := 1; i < len(opds); i++ {
			for _, p := range possibleAns {
				possibleAns = append(possibleAns, p*opds[i])
				possibleAns = append(possibleAns, p+opds[i])
				if possibleAns[len(possibleAns)-1] == res {
					sum += possibleAns[len(possibleAns)-1]
					break outer

				} else if possibleAns[len(possibleAns)-2] == res {
					sum += possibleAns[len(possibleAns)-2]
					break outer
				}
			}
		}
		// fmt.Printf("%v possibleAns: %v\n", res, possibleAns)
		_ = res
	}
	return sum
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)

	// oprands := map[int][]int{}

	sum := 0
	for sc.Scan() {
		l := sc.Text()
		sl := strings.Split(l, ": ")

		res, _ := strconv.Atoi(sl[0])
		opds := []int{}
		for _, opd := range strings.Split(sl[1], " ") {
			o, _ := strconv.Atoi(opd)
			opds = append(opds, o)
		}

		possibleAns := map[int]bool{opds[0]: true}
		for i := 1; i < len(opds); i++ {
			newPossibleAns := map[int]bool{}

			for p := range possibleAns {
				pd := p * opds[i]
				a := p + opds[i]
				j := joinNum(p, opds[i])

				if pd <= res {
					newPossibleAns[pd] = true
				}
				if a <= res {
					newPossibleAns[a] = true
				}
				if j <= res {
					newPossibleAns[j] = true
				}
			}
			possibleAns = newPossibleAns
		}

		if possibleAns[res] {
			sum += res
		}
		// return sum
	}
	return sum
}

func joinNum(a, b int) int {
	j := a
	m := 1
	for tmp := b; tmp > 0; tmp /= 10 {
		m *= 10
	}

	return j*m + b
}
