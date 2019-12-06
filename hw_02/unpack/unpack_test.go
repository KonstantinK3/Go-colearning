package unpack

import (
	"testing"
)

func TestUnpack(t *testing.T) {

	tables := []struct {
		input  string
		output string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"345", ""},
		{"a2b3n", "aabbbn"},
	}

	for tableNumber, table := range tables {

		unpacked := Unpack(table.input)
		if unpacked != table.output {
			t.Errorf("Case number %d, Input: %s, expected: %s, output: %s", tableNumber+1, table.input, table.output, unpacked)
		}

	}
}

// "github.com/stretchr/testify/assert"
// assert.Equal(t, Unpack(table.input), table.output, "unpacking "+table.input+" expecting "+table.output+" got "+Unpack(table.input))
// input := "a4bc2d5e"
// output := "aaaabccddddde"
// assert.Equal(t, Unpack(input), output, "unpacking "+input+" expecting "+output+" got "+Unpack(input))
