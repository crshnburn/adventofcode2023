package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMargin(t *testing.T) {
	input := [][2]int{{7, 9}, {15, 40}, {30, 200}}

	require.Equal(t, 288, FindMargin(input))
}
