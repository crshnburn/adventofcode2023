package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCoords(t *testing.T) {
	lines := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	require.Equal(t, 142, GetCoords(lines))
}

func TestGetCoords2(t *testing.T) {
	lines := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}

	require.Equal(t, 281, GetCoords2(lines))
}
