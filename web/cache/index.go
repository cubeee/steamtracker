package cache

import (
	"log"

	"github.com/cubeee/steamtracker/web/model"
	"github.com/cubeee/steamtracker/web/statistics"
)

type IndexCache struct {
	TrackedPlayers         uint
	CollectiveHoursTracked uint64
	GameStats24h           *[]model.GameStatistic
	GameStats7d            *[]model.GameStatistic
	GameStats              *[]model.GameStatistic
}

func LoadIndexCache() *IndexCache {
	games := 10

	var stats24h *[]model.GameStatistic
	var stats7d *[]model.GameStatistic
	var stats *[]model.GameStatistic

	loader := Loader{}
	loader.Add(func() {
		stats24h = statistics.GetGameStatistics24h(games)
	})
	loader.Add(func() {
		stats7d = statistics.GetGameStatistics7d(games)
	})
	loader.Add(func() {
		stats = statistics.GetGameStatistics(games)
	})

	elapsed := loader.LoadSync()
	log.Println("Index cache loaded in", elapsed)

	return &IndexCache{
		TrackedPlayers:         1000,   // TODO: real value from db
		CollectiveHoursTracked: 100000, // TODO: real value from db
		GameStats24h:           stats24h,
		GameStats7d:            stats7d,
		GameStats:              stats,
	}
}
