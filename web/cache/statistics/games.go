package statistics

import (
	"log"

	cacheLoader "github.com/cubeee/steamtracker/web/cache/loader"
	"github.com/cubeee/steamtracker/web/model"
	"github.com/cubeee/steamtracker/web/statistics"
)

type gameCache struct {
	GameStats24h []model.GameStatistic
	GameStats7d  []model.GameStatistic
	GameStats    []model.GameStatistic
}

var GameStatisticsCache = gameCache{}

func LoadGameStatisticsCache(games int64) {
	loader := cacheLoader.Loader{}
	loader.Add(func() {
		GameStatisticsCache.GameStats24h = *statistics.GetGameStatistics24h(games)
	})
	loader.Add(func() {
		GameStatisticsCache.GameStats7d = *statistics.GetGameStatistics7d(games)
	})
	loader.Add(func() {
		GameStatisticsCache.GameStats = *statistics.GetGameStatistics(games)
	})

	elapsed := loader.LoadSync()
	log.Println("Game statistics cache loaded in", elapsed)
}
