package cache

import (
	"log"

	"github.com/cubeee/steamtracker/web/model"
	"github.com/cubeee/steamtracker/web/statistics"
)

type IndexCache struct {
	TrackedPlayers         uint64
	CollectiveHoursTracked uint64
	GameStats24h           *[]model.GameStatistic
	GameStats7d            *[]model.GameStatistic
	GameStats              *[]model.GameStatistic
}

func LoadIndexCache(games int64) IndexCache {
	indexCache := IndexCache{}

	loader := Loader{}
	loader.Add(func() {
		indexCache.TrackedPlayers = statistics.GetTrackedPlayersCount()
	})
	loader.Add(func() {
		indexCache.CollectiveHoursTracked = statistics.GetCollectiveHoursTracked()
	})
	loader.Add(func() {
		indexCache.GameStats24h = statistics.GetGameStatistics24h(games)
	})
	loader.Add(func() {
		indexCache.GameStats7d = statistics.GetGameStatistics7d(games)
	})
	loader.Add(func() {
		indexCache.GameStats = statistics.GetGameStatistics(games)
	})

	elapsed := loader.LoadSync()
	log.Println("Index cache loaded in", elapsed)

	return indexCache
}
