package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input1 = []string{
	"7-F7-",
	".FJ|7",
	"SJLL7",
	"|F--J",
	"LJ.LJ",
}

var input = []string{
	"-L|F7",
	"7S-7|",
	"L|7||",
	"-L-J|",
	"L|-JF",
}

var findMiddle1 = []string{
	"...........",
	".S-------7.",
	".|F-----7|.",
	".||.....||.",
	".||.....||.",
	".|L-7.F-J|.",
	".|..|.|..|.",
	".L--J.L--J.",
	"...........",
}

var findMiddle2 = []string{
	"..........",
	".S------7.",
	".|F----7|.",
	".||....||.",
	".||....||.",
	".|L-7F-J|.",
	".|..||..|.",
	".L--JL--J.",
	"..........",
}

var findMiddle3 = []string{
	".F----7F7F7F7F-7....",
	".|F--7||||||||FJ....",
	".||.FJ||||||||L7....",
	"FJL7L7LJLJ||LJ.L-7..",
	"L--J.L7...LJS7F-7L7.",
	"....F-J..F7FJ|L7L7L7",
	"....L7.F7||L7|.L7L7|",
	".....|FJLJ|FJ|F7|.LJ",
	"....FJL-7.||.||||...",
	"....L---J.LJ.LJLJ...",
}

func TestCountSteps(t *testing.T) {

	steps, _ := CountSteps(input)
	require.Equal(t, 4, steps)
	steps, _ = CountSteps(input1)
	require.Equal(t, 8, steps)
}

func TestWallCounting(t *testing.T) {
	require.Equal(t, 4, FindMiddle(findMiddle1))

	require.Equal(t, 4, FindMiddle(findMiddle2))
	require.Equal(t, 8, FindMiddle(findMiddle3))
}
