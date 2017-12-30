package update

import (
	"log"
	"time"

	database "github.com/cubeee/steamtracker/shared/db"
	"github.com/cubeee/steamtracker/shared/db/paginator"
	"github.com/cubeee/steamtracker/shared/model"
	"github.com/cubeee/steamtracker/shared/steam"
	"github.com/cubeee/steamtracker/updater/task"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type SnapshotUpdater struct {
	task.Updater
}

type BatchItem struct {
	Identifier string
	Id         int64
}

func (updater SnapshotUpdater) Update() {
	log.Println("Updating player game snapshots...")
	resolver := &steam.Resolver{}

	options := &paginator.Options{}
	options.PageSize = uint64(viper.GetInt64("snapshot_update_batch_size"))
	p := paginator.NewPaginator(options, database.Db, &model.Player{})

	var batches int64
	var totalTime time.Duration

	var page uint64
	for page = 1; page <= p.Pages(); page++ {
		start := time.Now()
		batch := updater.getBatch(p, page)
		updater.UpdateBatch(resolver, batch)

		batches += 1
		elapsed := time.Since(start)
		totalTime += elapsed

		time.Sleep(2 * time.Second)
	}
	log.Println("Game snapshots updated:", batches, "batches completed in", totalTime.Seconds()+float64(batches*2),
		"seconds, average time per batch:", int64(totalTime/time.Millisecond)/batches, "ms")
}

func (updater SnapshotUpdater) UpdateBatch(resolver *steam.Resolver, items *[]BatchItem) {
	db := database.Db

	saveUpdate := func(identifier string, gameCount int) {
		db.
			Table(model.Player{}.TableName()).
			Where("identifier = ?", identifier).
			Updates(model.Player{
				GameCount:   gameCount,
				LastUpdated: time.Now().UTC(),
			})
	}

	for _, item := range *items {
		gameCount, games, err := resolver.GetGames(item.Identifier)
		if err != nil {
			log.Fatal(err)
			continue
		}

		if gameCount <= 0 || games == nil || len(*games) <= 0 {
			saveUpdate(item.Identifier, gameCount)
			continue
		}

		lastSnapshots := updater.getLastSnapshots(db, item.Id)
		for _, game := range *games {
			if game.Playtime == 0 {
				continue
			}

			db.Save(&game)

			hasSnapshots, lastSnapshot := updater.getLastSnapshot(lastSnapshots, &game)
			if !hasSnapshots || game.Playtime <= lastSnapshot.MinutesPlayed {
				continue
			}

			snapshot := model.NewSnapshot(item.Id, game.AppId, game.Playtime)
			db.Create(&snapshot)
		}

		saveUpdate(item.Identifier, gameCount)
	}
}

func (updater SnapshotUpdater) getLastSnapshot(snapshots *[]model.GameSnapshot, game *model.Game) (bool, *model.GameSnapshot) {
	if snapshots == nil {
		return false, &model.GameSnapshot{}
	}
	for _, snapshot := range *snapshots {
		if snapshot.GameId == game.AppId {
			return true, &snapshot
		}
	}
	return false, &model.GameSnapshot{}
}

func (updater SnapshotUpdater) getLastSnapshots(db *gorm.DB, playerId int64) *[]model.GameSnapshot {
	var lastSnapshots []model.GameSnapshot
	db.
		Table(model.GameSnapshot{}.TableName()).
		Where("player_id = ?", playerId).
		Order("game_id ASC, date DESC").
		Select("DISTINCT ON(game_id) *").
		Find(&lastSnapshots)
	return &lastSnapshots
}

func (updater SnapshotUpdater) getBatch(paginator *paginator.Paginator, page uint64) *[]BatchItem {
	var items []BatchItem
	paginator.PageCustom(page, items, func(db *gorm.DB, out interface{}) {
		db.Table(model.Player{}.TableName()).Select("id, identifier").Order("id asc").Find(&items)
	})
	return &items
}
