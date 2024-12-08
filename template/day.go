package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("INPUTFILEPATH")
	defer f.Close()
	fmt.Println(part1(f))
}

func part1(input io.Reader) int {
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		_ = sc.Text()
	}
	return 0
}

func part2(input io.Reader) int {
	sc := bufio.NewScanner(input)
	for sc.Scan() {
		_ = sc.Text()

	}
	return 0
}
