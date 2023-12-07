package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483"}

func TestFindWinnings(t *testing.T) {

	require.Equal(t, 6440, FindWinnings(input))
}

func TestFindWinningsWithJokers(t *testing.T) {
	require.Equal(t, 5905, FindWinningsWithJokers(input))
}
