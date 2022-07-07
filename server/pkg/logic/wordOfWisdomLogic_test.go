package logic

import (
	"testing"
	"wow/server/pkg/utils"
)

func TestPickRandomWordOfWisdom(t *testing.T) {
	wisdom := PickRandomWordOfWisdom()
	if !utils.IsStringElementInArray(wisdom, utils.WoW) {
		t.Errorf("%s is not in the array", wisdom)
	}
}