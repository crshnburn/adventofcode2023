package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestCountNextElements(t *testing.T) {

	part1, part2 := CountNextElements(input)
	require.Equal(t, 114, part1)
	require.Equal(t, 2, part2)
}
