package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func TestCountPaths(t *testing.T) {

	require.Equal(t, 374, CountPaths(input, 2))
	require.Equal(t, 1030, CountPaths(input, 10))
	require.Equal(t, 8410, CountPaths(input, 100))
}
