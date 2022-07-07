package logic

import (
	"testing"
)

func TestPrepareChallange(t *testing.T) {
	challange, answer := PrepareChallange()
	if challange != answer[0]*answer[1] {
		t.Errorf("challge is not solvable")
	}
}
