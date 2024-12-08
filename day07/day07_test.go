package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// func Test1(t *testing.T) {
// 	input := strings.NewReader(`190: 10 19
// 3267: 81 40 27
// 83: 17 5
// 156: 15 6
// 7290: 6 8 6 15
// 161011: 16 10 13
// 192: 17 8 14
// 21037: 9 7 18 13
// 292: 11 6 16 20`)

//		result := part1(input)
//		require.Equal(t, 3749, result)
//	}
func Test2(t *testing.T) {
	input := strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)

	result := part2(input)
	require.Equal(t, 11387, result)
}

func Test2Edge(t *testing.T) {
	input := strings.NewReader(`208098: 8 5 7 8 85 9 5 5 1 9 3 3
2064067023: 40 189 34 9 6 90 9 3 3 3
1096523: 9 9 57 13 3 54 516 7 2
12096608: 9 1 73 78 9 83 6 2 3 186
166183: 5 3 7 72 4 18 4 879 2 9 3
1191546920: 8 4 2 74 8 5 9 5 92 920 2
221791686495: 277 239 608 8 8 3 4 9 2
8357: 801 6 82 259 3`)

	result := part2(input)
	require.Equal(t, 0, result)
}
