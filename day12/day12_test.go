package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestCountPermutations(t *testing.T) {

	require.Equal(t, 21, CountPermutations(input))
}

func TestExpandedMap(t *testing.T) {
	require.Equal(t, 525152, ExpandedMap(input))
}
