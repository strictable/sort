package sort

import (
	"testing"
)

func TestInit(t *testing.T) {
	result := 2 + 3
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
