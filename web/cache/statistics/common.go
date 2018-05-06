package statistics

import (
	"log"

	cacheLoader "github.com/cubeee/steamtracker/web/cache/loader"
	"github.com/cubeee/steamtracker/web/statistics"
)

type commonStatistics struct {
	TrackedPlayers         uint64
	CollectiveHoursTracked uint64
}

var CommonStatisticsCache = commonStatistics{}

func LoadCommonStatisticsCache() {
	loader := cacheLoader.Loader{}
	loader.Add(func() {
		CommonStatisticsCache.TrackedPlayers = statistics.GetTrackedPlayersCount()
	})
	loader.Add(func() {
		CommonStatisticsCache.CollectiveHoursTracked = statistics.GetCollectiveHoursTracked()
	})

	elapsed := loader.LoadSync()
	log.Println("Common statistics cache loaded in", elapsed)
}
