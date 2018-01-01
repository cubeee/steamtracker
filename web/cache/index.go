package cache

import (
	sharedModel "github.com/cubeee/steamtracker/shared/model"
	"github.com/cubeee/steamtracker/web/model"
)

type IndexCache struct {
	TrackedPlayers         uint
	CollectiveHoursTracked uint64
	GameStats24h           *[]model.GameStatistic
	GameStats7d            *[]model.GameStatistic
	GameStats              *[]model.GameStatistic
}

func PreloadIndexCache() *IndexCache {
	return &IndexCache{
		TrackedPlayers:         1000,   // TODO: real value from db
		CollectiveHoursTracked: 100000, // TODO: real value from db
		GameStats: &[]model.GameStatistic{ // TODO: real stats
			{999, sharedModel.Game{
				AppId: 578080,
				Name:  "PUBG herp derp",
				Icon:  "93d896e7d7a42ae35c1d77239430e1d90bc82cae",
			}},
		},
	}
}
