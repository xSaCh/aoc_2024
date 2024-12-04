package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./day03/input.txt")
	defer f.Close()
	fmt.Println(part2(f))
}

func part1(input io.Reader) int {
	rg := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	memData, _ := io.ReadAll(input)
	matches := rg.FindAllString(string(memData), -1)

	sum := 0
	for _, v := range matches {
		sum += parseMul(v)
	}
	return sum
}

func part2(input io.Reader) int {
	rg := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|do\(\)|don't\(\)`)
	memData, _ := io.ReadAll(input)
	matches := rg.FindAllString(string(memData), -1)
	isEnable := true
	sum := 0
	for _, v := range matches {
		if v == "do()" {
			isEnable = true
		} else if v == "don't()" {
			isEnable = false

		} else if isEnable {
			// fmt.Printf("v: %v\n", v)
			sum += parseMul(v)
		}

	}
	return sum
}

func parseMul(token string) int {
	vals := strings.Split(token[4:len(token)-1], ",")
	a, _ := strconv.Atoi(vals[0])
	b, _ := strconv.Atoi(vals[1])
	return a * b
}
