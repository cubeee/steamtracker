package main

import (
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/cubeee/steamtracker/shared/config"
	"github.com/cubeee/steamtracker/shared/db"
	"github.com/cubeee/steamtracker/web/bootstrap"
	"github.com/cubeee/steamtracker/web/cache"
	"github.com/cubeee/steamtracker/web/routes"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/spf13/viper"
)

const (
	AppName       = "SteamTracker"
	MigrationsDir = "file://./resources/migrations"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	bootStart := time.Now()

	env := flag.String("env", "dev", "Program execution environment")
	flag.Parse()

	log.Println("SteamTracker Web starting...")
	log.Println("Environment:", *env)
	config.ReadConfig("web", env)

	connectDetails := &db.ConnectDetails{
		Host:       config.GetString("database.host"),
		Db:         config.GetString("database.name"),
		User:       config.GetString("database.username"),
		Pass:       config.GetString("database.password"),
		Additional: "sslmode=disable",
	}
	if err := db.ConnectPostgres(connectDetails); err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()
	migrateDatabase(connectDetails)

	globalContext := &bootstrap.GlobalCtx{
		GoogleAnalyticsId: config.GetString("google-analytics-id"),
	}

	preloadCache()

	app := bootstrap.New(AppName)
	app.Bootstrap(env, globalContext)
	app.Configure(routes.Configure)

	bootElapsed := time.Since(bootStart)
	log.Println("SteamTracker Web started in", bootElapsed)

	app.Listen(config.GetString("server.addr"))
}

func migrateDatabase(connectDetails *db.ConnectDetails) {
	log.Println("Running migrations...")
	driver, err := postgres.WithInstance(db.Db.DB(), &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(MigrationsDir, connectDetails.Db, driver)
	if err != nil {
		log.Panic(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("Database schema up to date, no migrations needed")
		} else {
			log.Panic(err)
		}
	} else {
		log.Println("Migrations executed successfully")
	}
}

func preloadCache() {
	cache.GlobalCache = &cache.Cache{}
	cache.GlobalCache.SetIndexCache(cache.LoadIndexCache())
}
