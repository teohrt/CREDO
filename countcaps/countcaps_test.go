package countcaps

import (
	"testing"
)

func TestCountCaps(t *testing.T) {
	s := "Trace"
	num := CountCaps(s)

	if num != 1 {
		t.Errorf("Expected 1. Got: %d", num)
	}

	s = "ALLCAPSMAYBE"
	num = CountCaps(s)

	if num != 12 {
		t.Errorf("Expected 12. Got: %d", num)
	}

	s = ""
	num = CountCaps(s)

	if num != 0 {
		t.Errorf("Expected 0. Got: %d", num)
	}

	s = "999"
	num = CountCaps(s)

	if num != 0 {
		t.Errorf("Expected 0. Got: %d", num)
	}

	s = ".&*"
	num = CountCaps(s)

	if num != 0 {
		t.Errorf("Expected 0. Got: %d", num)
	}
}
