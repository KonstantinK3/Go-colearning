package wordcount

import (
	"testing"
)

func TestWordcount(t *testing.T) {

	tables := []struct {
		input  string
		output []string
	}{
		{"../test.txt", []string{"шесть", "пять", "четыре", "три", "два"}},
	}

	for tableNumber, table := range tables {

		result := Wordcount(table.input)
		if result != nil {
			t.Errorf("Case number %d, Input: %s, expected: %s, output: %s", tableNumber+1, table.input, table.output, result)
		}

	}
}
