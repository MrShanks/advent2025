package p2

import (
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("Solution", func(t *testing.T) {
		got := Solve("../input_test.txt")
		expected := 4174379265

		if got != expected {
			t.Errorf("Expected: %d\nGot %d\n", expected, got)
		}
	})
}
