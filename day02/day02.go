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
	f, _ := os.Open("./day02/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	reports := make([][]int, 0)
	rc := 0
	for sc.Scan() {
		line := sc.Text()
		rpt := strings.Split(line, " ")

		reports = append(reports, make([]int, len(rpt)))
		for i, r := range rpt {
			ri, _ := strconv.Atoi(r)
			reports[rc][i] = ri
		}
		rc++
	}
	cnt := 0
	for _, r := range reports {
		isDecreasing := false
		if (r[0] - r[1]) > 0 {
			isDecreasing = true
		}
		safeState := true
		for i := 0; i < len(r)-1; i++ {
			diff := r[i] - r[i+1]
			if !isDecreasing {
				diff = -diff
			}

			if diff < 1 || diff > 3 {
				safeState = false
				break
			}
		}
		if safeState {
			cnt++
		}
	}
	return cnt
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	reports := make([][]int, 0)
	rc := 0
	for sc.Scan() {
		line := sc.Text()
		rpt := strings.Split(line, " ")

		reports = append(reports, make([]int, len(rpt)))
		for i, r := range rpt {
			ri, _ := strconv.Atoi(r)
			reports[rc][i] = ri
		}
		rc++
	}
	cnt := 0
	for _, rr := range reports {
		subReports := make([][]int, 0)
		for i := range rr {
			subReports = append(subReports, make([]int, 0, len(rr)-1))
			for j := 0; j < len(rr); j++ {
				if j == i {
					continue
				}
				subReports[i] = append(subReports[i], rr[j])
			}
		}
		st := 0
		for _, r := range subReports {

			isDecreasing := false
			safeState := true
			if (r[0] - r[1]) > 0 {
				isDecreasing = true
			}
			for i := 0; i < len(r)-1; i++ {
				diff := r[i] - r[i+1]
				if !isDecreasing {
					diff = -diff
				}

				if diff < 1 || diff > 3 {
					safeState = false
					break
				}
			}
			if safeState {
				st++
				break
			}
		}
		if st > 0 {
			cnt++
		}
	}
	return cnt
}
