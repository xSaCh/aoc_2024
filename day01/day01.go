package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day01/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	left := make([]int, 0)
	right := make([]int, 0)

	for sc.Scan() {
		line := sc.Text()
		val := strings.Split(line, "   ")

		v1, _ := strconv.Atoi(val[0])
		v2, _ := strconv.Atoi(val[1])
		left = append(left, v1)
		right = append(right, v2)
	}

	slices.Sort(left)
	slices.Sort(right)
	diff := 0.0
	for i := range len(left) {
		diff += math.Abs(float64(left[i] - right[i]))
	}
	return int(diff)
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	left := make([]int, 0)
	right := make([]int, 0)

	for sc.Scan() {
		line := sc.Text()
		val := strings.Split(line, "   ")

		v1, _ := strconv.Atoi(val[0])
		v2, _ := strconv.Atoi(val[1])
		left = append(left, v1)
		right = append(right, v2)
	}

	freq := map[int]int{}
	for _, i := range right {
		v, ex := freq[i]
		if !ex {
			freq[i] = 1
		} else {
			freq[i] = v + 1
		}
	}

	score := 0
	for _, v := range left {
		f, _ := freq[v]
		score += f * v
	}
	return int(score)
}
