package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1SimpleAntiNode(t *testing.T) {
	input := strings.NewReader(`..........
..........
..........
....a.....
........a.
.....a....
..........
......A...
..........
..........
`)

	result := part1(input)
	require.Equal(t, 4, result)
}

func Test1(t *testing.T) {
	input := strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

	result := part1(input)
	require.Equal(t, 14, result)
}

func Test2SimpleAntiNode(t *testing.T) {
	input := strings.NewReader(`T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`)

	result := part2(input)
	require.Equal(t, 9, result)
}

func Test2(t *testing.T) {
	input := strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)

	result := part2(input)
	require.Equal(t, 34, result)
}