package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test1(t *testing.T) {
	input := strings.NewReader(``)

	result := part1(input)
	require.Equal(t, 0, result)
}
func Test2(t *testing.T) {
	input := strings.NewReader(``)

	result := part2(input)
	require.Equal(t, 0, result)
}
