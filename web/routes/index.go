package routes

import (
	statisticsCache "github.com/cubeee/steamtracker/web/cache/statistics"
	"github.com/kataras/iris"
)

func GetIndexHandler(ctx iris.Context) {
	ctx.ViewData("tracked_players", statisticsCache.CommonStatisticsCache.TrackedPlayers)
	ctx.ViewData("collective_hours_tracked", statisticsCache.CommonStatisticsCache.CollectiveHoursTracked)
	ctx.ViewData("game_stats_24h", statisticsCache.GameStatisticsCache.GameStats24h)
	ctx.ViewData("game_stats_7d", statisticsCache.GameStatisticsCache.GameStats7d)
	ctx.ViewData("game_stats", statisticsCache.GameStatisticsCache.GameStats)
	ctx.View("index.tpl")
}
