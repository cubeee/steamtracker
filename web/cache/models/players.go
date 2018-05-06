package models

import (
	database "github.com/cubeee/steamtracker/shared/db"
	"github.com/cubeee/steamtracker/shared/model"
)

var playerCache = map[int64]*model.Player{}

func GetCachedPlayer(id int64) (*model.Player, bool) {
	player, err := playerCache[id]
	return player, err
}

func GetCachedPlayerOrLoad(id int64) model.Player {
	cached, ok := GetCachedPlayer(id)
	var player model.Player
	if !ok {
		database.Db.Where("id = ?", id).First(&player)
		CachePlayer(&player)
		return player
	}
	return *cached
}

func CachePlayer(player *model.Player) {
	playerCache[player.Id] = player
}
