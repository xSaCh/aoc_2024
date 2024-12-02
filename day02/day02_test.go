package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)

	result := part1(input)
	require.Equal(t, 2, result)
}

func Test2(t *testing.T) {
	input := strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`)
	_ = input
	result := part2(input)
	require.Equal(t, 4, result)
}

func Test2EdgeCase(t *testing.T) {
	input := strings.NewReader(`48 46 47 49 51 54 56
1 1 2 3 4 5
1 2 3 4 5 5
5 1 2 3 4 5
1 4 3 2 1
1 6 7 8 9
1 2 3 4 3
9 8 7 6 7
7 10 8 10 11
29 28 27 25 26 25 22 20`)

	result := part2(input)
	require.Equal(t, 10, result)
}
