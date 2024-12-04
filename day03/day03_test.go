package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(`xmul(2,4)%
&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`)

	result := part1(input)
	require.Equal(t, 161, result)
}

func Test2(t *testing.T) {
	input := strings.NewReader(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5)))`)

	result := part2(input)
	require.Equal(t, 48, result)
}
