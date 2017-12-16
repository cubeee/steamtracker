package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/cubeee/steamtracker/shared/db"
	"github.com/cubeee/steamtracker/updater/task/update"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

var (
	cronScheduler *cron.Cron
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	log.Println("SteamTracker Updater starting...")

	viper.SetConfigName("updater-config")
	viper.AddConfigPath("./resources/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	connectDetails := db.ConnectDetails{
		Host:       os.Getenv("DB_HOST"),
		Db:         os.Getenv("DB"),
		User:       os.Getenv("DB_USER"),
		Pass:       os.Getenv("DB_PASS"),
		Additional: "sslmode=disable",
	}
	if err := db.ConnectPostgres(connectDetails); err != nil {
		panic(err)
	}
	defer db.Db.Close()

	scheduleTasks()
	listenSignals()
}

func scheduleTasks() {
	cronScheduler = cron.New()

	if viper.GetBool("enable_profile_updates") {
		cronScheduler.AddFunc(viper.GetString("profile_update_cron"), update.ProfileUpdater{}.Update)
	}

	go func() {
		cronScheduler.Start()
	}()
}

func listenSignals() {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			log.Println("Interrupt signal received, stopping")
			cronScheduler.Stop()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
