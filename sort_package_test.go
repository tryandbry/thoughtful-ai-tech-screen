package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPackage(t *testing.T) {
	testCases := []struct {
		width    int
		height   int
		length   int
		mass     int
		expected string
	}{
		{1, 1, 1, 1, "STANDARD"},
		{150, 1, 1, 1, "SPECIAL"},
		{149, 149, 149, 1, "SPECIAL"},
		{1, 1, 1, 20, "SPECIAL"},
		{149, 149, 149, 20, "REJECTED"},
		{0, 149, 149, 20, "REJECTED"},
	}
	for _, tc := range testCases {
		t.Logf("dimensions: %dx%dx%d, mass: %dkg\n", tc.width, tc.height, tc.length, tc.mass)
		assert.Equal(t, tc.expected, sortPackage(tc.width, tc.height, tc.length, tc.mass))
	}
}
