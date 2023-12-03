package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindParts(t *testing.T) {
	input := []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}

	require.Equal(t, 4361, FindParts(input))
}

func TestFindGears(t *testing.T) {
	input := []string{"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598.."}

	require.Equal(t, 467835, FindGears(input))
}
