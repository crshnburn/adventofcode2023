package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

func TestCountReflections(t *testing.T) {

	require.Equal(t, 405, CountReflections(input))
}

func TestCountReflectionsWithSmudge(t *testing.T) {
	require.Equal(t, 400, CountReflectionsWithSmudge(input))
}
