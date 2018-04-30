package statistics

import (
	"time"

	database "github.com/cubeee/steamtracker/shared/db"
	sharedModels "github.com/cubeee/steamtracker/shared/model"
	"github.com/cubeee/steamtracker/web/model"
)

var gameCache = map[int]*sharedModels.Game{}

type statistic struct {
	GameId        int   `gorm:"column:game_id"`
	MinutesPlayed int64 `gorm:"column:sum"`
}

func GetGameStatistics24h(limit int) *[]model.GameStatistic {
	now := time.Now()
	day := now.Add(-24 * time.Hour)
	return getOverallGameStatistics(day, now, limit)
}

func GetGameStatistics7d(limit int) *[]model.GameStatistic {
	now := time.Now()
	week := now.Add(-(24 * 7) * time.Hour)
	return getOverallGameStatistics(week, now, limit)
}

func GetGameStatistics(limit int) *[]model.GameStatistic {
	now := time.Now()
	start := time.Unix(0, 0)
	return getOverallGameStatistics(start, now, limit)
}

func getOverallGameStatistics(from, to time.Time, limit int) *[]model.GameStatistic {
	db := database.Db
	var statistics []statistic
	var gameStats []model.GameStatistic
	db.Raw("SELECT * FROM all_games_minutes_tracked(?, ?) AS f(game_id BIGINT, sum BIGINT) LIMIT ?",
		from, to, limit).Scan(&statistics)
	for _, statistic := range statistics {
		game := getGameFromCache(statistic.GameId)

		gameStats = append(gameStats, model.GameStatistic{Game: game, MinutesPlayed: statistic.MinutesPlayed})
	}
	return &gameStats
}

func getGameFromCache(id int) sharedModels.Game {
	cached, ok := gameCache[id]
	var game sharedModels.Game
	if !ok {
		database.Db.Where("app_id = ?", id).First(&game)
		gameCache[id] = &game
		return game
	}
	return *cached
}
