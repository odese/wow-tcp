package logic

import (
	math "math/rand"
	"wow/server/pkg/utils"
)

func PickRandomWordOfWisdom() string {
	randomIndex := math.Intn(len(utils.WoW))
	pick := utils.WoW[randomIndex]
	return pick
}
