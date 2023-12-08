package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input1 = []string{"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)"}

var input2 = []string{"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)"}

var ghostInput = []string{"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)"}

func TestCountSteps(t *testing.T) {

	require.Equal(t, 2, CountSteps(input1))
	require.Equal(t, 6, CountSteps(input2))
}

func TestGhostSteps(t *testing.T) {
	require.Equal(t, 6, GhostSteps(ghostInput))
}
