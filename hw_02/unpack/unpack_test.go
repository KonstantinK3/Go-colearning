package unpack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {

	input := "a4bc2d5e"
	output := "aaaabccddddde"
	assert.Equal(t, Unpack(input), output, "unpacking "+input+" expecting "+output+" got "+Unpack(input))

	input = "abcd"
	output = "abcd"
	assert.Equal(t, Unpack(input), output, "unpacking "+input+" expecting "+output+" got "+Unpack(input))

	input = "45"
	output = ""
	assert.Equal(t, Unpack(input), output, "unpacking "+input+" expecting "+output+" got "+Unpack(input))

}
