package models

import (
	database "github.com/cubeee/steamtracker/shared/db"
	"github.com/cubeee/steamtracker/shared/model"
)

var gameCache = map[int64]*model.Game{}

func GetCachedGame(id int64) (*model.Game, bool) {
	game, err := gameCache[id]
	return game, err
}

func GetCachedGameOrLoad(id int64) model.Game {
	cached, ok := GetCachedGame(id)
	var game model.Game
	if !ok {
		database.Db.Where("app_id = ?", id).First(&game)
		CacheGame(&game)
		return game
	}
	return *cached
}

func CacheGame(game *model.Game) {
	gameCache[game.AppId] = game
}
