package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
    act1 := Round(float64(10.124), float64(2))
    act2 := Round(float64(10.125), float64(2))
    assert.Equal(t, float64(10.12), act1)
    assert.Equal(t, float64(10.13), act2)
}

func TestRoundUp(t *testing.T) {
    act1 := RoundUp(float64(10.124), float64(2))
    act2 := RoundUp(float64(10.125), float64(2))
    assert.Equal(t, float64(10.13), act1)
    assert.Equal(t, float64(10.13), act2)
}

func TestRoundDown(t *testing.T) {
    act1 := RoundDown(float64(10.124), float64(2))
    act2 := RoundDown(float64(10.125), float64(2))
    assert.Equal(t, float64(10.12), act1)
    assert.Equal(t, float64(10.12), act2)
}