package logic

import "testing"

func TestPrepareProofMessageForServer(t *testing.T) {
	divisors := make([]uint64, 0)
	divisors = append(divisors, 1)
	divisors = append(divisors, 2)

	expected := "1 2 "

	got := prepareProofMessageForServer(divisors)
	
	if expected != got {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
