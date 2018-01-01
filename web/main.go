package main

import (
	"flag"
	"log"
	"runtime"

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
	env := flag.String("env", "dev", "Program execution environment")
	flag.Parse()

	log.Println("SteamTracker Web starting...")
	config.ReadConfig("web", env)

	connectDetails := &db.ConnectDetails{
		Host:       viper.GetString("database.host"),
		Db:         viper.GetString("database.name"),
		User:       viper.GetString("database.username"),
		Pass:       viper.GetString("database.password"),
		Additional: "sslmode=disable",
	}
	if err := db.ConnectPostgres(connectDetails); err != nil {
		log.Fatal(err)
	}
	defer db.Db.Close()
	migrateDatabase(connectDetails)

	globalContext := &bootstrap.GlobalCtx{
		GoogleAnalyticsId: viper.GetString("google-analytics-id"),
	}

	preloadCache()

	app := bootstrap.New(AppName)
	app.Bootstrap(env, globalContext)
	app.Configure(routes.Configure)
	app.Listen(viper.GetString("server.addr"))
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
		log.Panic(err)
	} else {
		log.Println("Migrations executed successfully")
	}
}

func preloadCache() {
	cache.GlobalCache = &cache.Cache{}
	cache.GlobalCache.SetIndexCache(cache.PreloadIndexCache())
}
