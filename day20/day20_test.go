package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"broadcaster -> a",
	"%a -> inv, con",
	"&inv -> b",
	"%b -> con",
	"&con -> output",
}

func TestCountPulses(t *testing.T) {

	require.Equal(t, 11687500, CountPulses(input))
}
