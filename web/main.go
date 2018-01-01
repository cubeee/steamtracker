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
	"github.com/spf13/viper"
)

const (
	AppName = "SteamTracker"
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

	globalContext := &bootstrap.GlobalCtx{
		GoogleAnalyticsId: viper.GetString("google-analytics-id"),
	}

	preloadCache()

	app := bootstrap.New(AppName)
	app.Bootstrap(env, globalContext)
	app.Configure(routes.Configure)
	app.Listen(viper.GetString("server.addr"))
}

func preloadCache() {
	cache.GlobalCache = &cache.Cache{}
	cache.GlobalCache.SetIndexCache(cache.PreloadIndexCache())
}
