package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day05/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)

	_, revRules := parseOrderingRules(sc)
	pages := parsePages(sc)

	midSum := 0
	for _, p := range pages {
		if isOrdered(p, revRules) {
			midSum += p[len(p)/2]
		}
	}

	return midSum
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)

	rules, revRules := parseOrderingRules(sc)
	pages := parsePages(sc)

	midSum := 0
	for _, p := range pages {
		if !isOrdered(p, revRules) {
			slices.SortFunc(p, func(x, y int) int {
				if slices.Contains(rules[x], y) {
					return -1
				}
				if slices.Contains(rules[y], x) {
					return 1
				}
				return 0
			})
			midSum += p[len(p)/2]
		}
	}

	return midSum
}

func isOrdered(page []int, revRules map[int][]int) bool {
	disAllowed := map[int]bool{}

	for _, p := range page {
		if _, ok := disAllowed[p]; ok {
			return false
		}

		for _, r := range revRules[p] {
			disAllowed[r] = true
		}
		// fmt.Printf("%d %v\n", p, disAllowed)

	}
	return true
}

func parseOrderingRules(sc *bufio.Scanner) (map[int][]int, map[int][]int) {
	rules := make(map[int][]int)
	revRules := make(map[int][]int)
	for sc.Scan() {
		l := sc.Text()
		if l == "" {
			break
		}
		sp := strings.Split(l, "|")
		a, _ := strconv.Atoi(sp[0])
		b, _ := strconv.Atoi(sp[1])

		if _, ok := rules[a]; !ok {
			rules[a] = make([]int, 0)
		}

		rules[a] = append(rules[a], b)

		if _, ok := revRules[b]; !ok {
			revRules[b] = make([]int, 0)
		}

		revRules[b] = append(revRules[b], a)
	}
	return rules, revRules
}

func parsePages(sc *bufio.Scanner) [][]int {
	pages := make([][]int, 0)
	_ = pages
	for sc.Scan() {
		l := sc.Text()
		sp := strings.Split(l, ",")
		updates := make([]int, 0, len(sp))
		for _, v := range sp {
			p, _ := strconv.Atoi(v)
			updates = append(updates, p)
		}
		pages = append(pages, updates)
	}
	return pages
}
