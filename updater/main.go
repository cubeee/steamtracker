package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/cubeee/steamtracker/shared/config"
	"github.com/cubeee/steamtracker/shared/db"
	"github.com/cubeee/steamtracker/updater/task/update"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
	metrics "github.com/tevjef/go-runtime-metrics"
)

var (
	cronScheduler *cron.Cron
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	env := flag.String("env", "dev", "Program execution environment")
	flag.Parse()

	log.Println("SteamTracker Updater starting...")
	config.ReadConfig("updater", env)

	connectDetails := &db.ConnectDetails{
		Host:       os.Getenv("DB_HOST"),
		Db:         os.Getenv("DB"),
		User:       os.Getenv("DB_USER"),
		Pass:       os.Getenv("DB_PASS"),
		Additional: "sslmode=disable",
	}
	if err := db.ConnectPostgres(connectDetails); err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()

	if viper.GetBool("enable_metrics_collecting") {
		metrics.DefaultConfig.CollectionInterval = time.Second
		metrics.DefaultConfig.BatchInterval = time.Second * 5
		metrics.DefaultConfig.Host = viper.GetString("metrics_host")
		metrics.DefaultConfig.Database = viper.GetString("metrics_database")
		if err := metrics.RunCollector(metrics.DefaultConfig); err != nil {
			log.Println("Metrics collection setup failed:", err)
		}
	}

	scheduleTasks()
	listenSignals()
}

func scheduleTasks() {
	log.Println("Scheduling tasks...")
	cronScheduler = cron.New()

	if viper.GetBool("enable_profile_updates") {
		cronScheduler.AddFunc(viper.GetString("profile_update_cron"), update.ProfileUpdater{}.Update)
	}

	if viper.GetBool("enable_snapshot_updates") {
		cronScheduler.AddFunc(viper.GetString("snapshot_update_cron"), update.SnapshotUpdater{}.Update)
	}

	go func() {
		cronScheduler.Start()
	}()
}

func listenSignals() {
	log.Println("Ready")
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
