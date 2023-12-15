package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input = []string{
	"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
}

func TestCreateHash(t *testing.T) {

	require.Equal(t, 1320, CreateHash(input))
}

func TestCalcFocusingPower(t *testing.T) {
	require.Equal(t, 145, CalcFocusingPower(input))
}
