package update

import (
	"log"
	"time"

	"github.com/cubeee/steamtracker/shared/db"
	"github.com/cubeee/steamtracker/shared/db/paginator"
	"github.com/cubeee/steamtracker/shared/model"
	"github.com/cubeee/steamtracker/shared/steam"
	"github.com/cubeee/steamtracker/updater/task"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type ProfileUpdater struct {
	task.Updater
}

func (updater ProfileUpdater) Update() {
	log.Println("Updating player profiles")
	resolver := steam.Resolver{}

	options := paginator.Options{}
	options.PageSize = viper.GetInt64("profile_update_batch_size")
	p := paginator.NewPaginator(options, db.Db, &model.Player{})

	var batches int64
	var totalTime time.Duration

	var page int64
	for page = 1; page <= 1; /*p.Pages()*/ page++ {
		start := time.Now()
		batch := updater.getIdentifierBatch(p, 1)
		updater.UpdateBatch(resolver, batch)

		batches += 1
		elapsed := time.Since(start)
		totalTime += elapsed

		time.Sleep(2 * time.Second)
	}
	log.Println(batches, "batches completed in", totalTime.Seconds()+float64(batches*2), "seconds, average time per batch:", int64(totalTime/time.Millisecond)/batches, "ms")
}

func (updater ProfileUpdater) UpdateBatch(resolver steam.Resolver, identifiers []string) {
	profiles, err := resolver.GetProfile(identifiers)
	if err != nil {
		log.Fatal(err)
	}

	tx := db.Db.Begin()
	for _, profile := range profiles {
		tx.
			Table(model.Player{}.TableName()).
			Where("identifier = ?", profile.Identifier).
			Updates(model.Player{
				Name:         profile.Name,
				Avatar:       profile.Avatar,
				AvatarMedium: profile.AvatarMedium,
				AvatarFull:   profile.AvatarFull,
				CountryCode:  profile.CountryCode,
			})
	}
	tx.Commit()
}

func (ProfileUpdater) getIdentifierBatch(paginator paginator.Paginator, page uint) []string {
	var identifiers []string
	paginator.PageCustom(1, identifiers, func(db *gorm.DB, out interface{}) {
		db.Table(model.Player{}.TableName()).Pluck("identifier", &identifiers)
	})
	return identifiers
}
