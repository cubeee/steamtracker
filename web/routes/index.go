package routes

import (
	"github.com/cubeee/steamtracker/web/cache"
	"github.com/kataras/iris"
)

func GetIndexHandler(ctx iris.Context) {
	indexCache := cache.GlobalCache.GetIndexCache()

	ctx.ViewData("tracked_players", indexCache.TrackedPlayers)
	ctx.ViewData("collective_hours_tracked", indexCache.CollectiveHoursTracked)
	ctx.ViewData("game_stats_24h", indexCache.GameStats24h)
	ctx.ViewData("game_stats_7d", indexCache.GameStats7d)
	ctx.ViewData("game_stats", indexCache.GameStats)
	ctx.View("index.tpl")
}
