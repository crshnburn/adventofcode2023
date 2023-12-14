package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

func TestCalculateLoad(t *testing.T) {

	require.Equal(t, 136, CalculateLoad(input))
}

func TestCalculateLoad2(t *testing.T) {
	require.Equal(t, 64, CalculateLoad2(input))
}
