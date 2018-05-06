package statistics

import (
	"time"

	database "github.com/cubeee/steamtracker/shared/db"
	sharedModels "github.com/cubeee/steamtracker/shared/model"
	modelCache "github.com/cubeee/steamtracker/web/cache/models"
	"github.com/cubeee/steamtracker/web/model"
)

type statistic struct {
	GameId        int64 `gorm:"column:game_id"`
	MinutesPlayed int64 `gorm:"column:minutes"`
}

func GetTrackedPlayersCount() uint64 {
	var count uint64
	database.Db.Table(sharedModels.Player{}.TableName()).Count(&count)
	return count
}

func GetCollectiveHoursTracked() uint64 {
	db := database.Db
	var sum []uint64
	db.Raw("SELECT * FROM total_minutes_played").Pluck("minutes", &sum)
	return sum[0]
}

func GetGameStatistics24h(limit int64) *[]model.GameStatistic {
	now := time.Now()
	day := now.Add(-24 * time.Hour)
	return getOverallGameStatistics(day, now, limit)
}

func GetGameStatistics7d(limit int64) *[]model.GameStatistic {
	now := time.Now()
	week := now.Add(-(24 * 7) * time.Hour)
	return getOverallGameStatistics(week, now, limit)
}

func GetGameStatistics(limit int64) *[]model.GameStatistic {
	now := time.Now()
	start := time.Unix(0, 0)
	return getOverallGameStatistics(start, now, limit)
}

func getOverallGameStatistics(from, to time.Time, limit int64) *[]model.GameStatistic {
	db := database.Db
	var statistics []statistic
	var gameStats []model.GameStatistic
	db.Raw("SELECT * FROM all_games_minutes_tracked(?, ?) AS f(game_id BIGINT, minutes BIGINT) LIMIT ?",
		from, to, limit).Scan(&statistics)
	for _, statistic := range statistics {
		game := modelCache.GetCachedGameOrLoad(statistic.GameId)

		gameStats = append(gameStats, model.GameStatistic{Game: game, MinutesPlayed: statistic.MinutesPlayed})
	}
	return &gameStats
}
